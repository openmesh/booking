package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/openmesh/booking"
	entbooking "github.com/openmesh/booking/ent/booking"
	"github.com/openmesh/booking/ent/resource"
)

type bookingService struct {
	client *Client
}

// NewBookingService constructs a new instance of a booking.BookingService using
// ent as its persistence layer.
func NewBookingService(client *Client) *bookingService {
	return &bookingService{client}
}

// Retrieves a single booking by ID along with the associated resource and
// metadata. Returns ENOTFOUND if booking does not exist or user does not have
// permission to view it.
func (s *bookingService) FindBookingByID(
	ctx context.Context,
	req booking.FindBookingByIDRequest,
) booking.FindBookingByIDResponse {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return booking.FindBookingByIDResponse{
			Err: fmt.Errorf("failed to start transaction: %w", err),
		}
	}

	b, err := findBookingByID(ctx, tx, req.ID)
	var nfe *NotFoundError
	if errors.As(err, &nfe) {
		return booking.FindBookingByIDResponse{
			Err: booking.WrapNotFoundError("booking"),
		}
	}
	if err != nil {
		return booking.FindBookingByIDResponse{
			Err: fmt.Errorf("failed to find booking by id: %w", err),
		}
	}
	if err := b.attachEdges(ctx); err != nil {
		return booking.FindBookingByIDResponse{
			Err: fmt.Errorf("failed to attach edges to booking: %w", err),
		}
	}

	return booking.FindBookingByIDResponse{Booking: b.toModel()}
}

func findBookingByID(ctx context.Context, tx *Tx, id int) (*Booking, error) {
	orgID := booking.OrganizationIDFromContext(ctx)
	return tx.Booking.
		Query().
		Where(
			entbooking.ID(id),
			entbooking.HasResourceWith(resource.OrganizationIdEQ(orgID)),
		).
		First(ctx)
}

// Retreives a list of bookings based on a filter. Only returns bookings that
// are accessible to the user. Also returns a count of total matching bookings
// which may be different from the number of returned bookings if the "Limit"
// field is set.
func (s *bookingService) FindBookings(
	ctx context.Context,
	req booking.FindBookingsRequest,
) booking.FindBookingsResponse {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return booking.FindBookingsResponse{
			Err: fmt.Errorf("failed to start transaction: %w", err),
		}
	}
	orgID := booking.OrganizationIDFromContext(ctx)
	query := tx.Booking.Query().
		Where(entbooking.HasResourceWith(resource.OrganizationId(orgID)))
	if req.ID != nil {
		query = query.Where(entbooking.ID(*req.ID))
	}
	if req.ResourceID != nil {
		query = query.Where(entbooking.ResourceId(*req.ResourceID))
	}
	if req.Status != nil {
		query = query.Where(entbooking.Status(*req.Status))
	}
	if req.StartTimeAfter != nil {
		query = query.Where(entbooking.StartTimeGTE(*req.StartTimeAfter))
	}
	if req.EndTimeBefore != nil {
		query = query.Where(entbooking.EndTimeLTE(*req.EndTimeBefore))
	}
	totalItems, err := query.Count(ctx)
	if err != nil {
		return booking.FindBookingsResponse{
			Err: fmt.Errorf("failed to count bookings: %w", err),
		}
	}
	query = query.Offset(req.Offset)
	if req.Limit == 0 {
		query = query.Limit(10)
	} else {
		query = query.Limit(req.Limit)
	}
	bookings, err := query.
		WithMetadata().
		WithResource().
		All(ctx)
	if err != nil {
		return booking.FindBookingsResponse{
			Err: fmt.Errorf("failed to query bookings: %w", err),
		}
	}

	return booking.FindBookingsResponse{
		Bookings:   Bookings(bookings).toModels(),
		TotalItems: totalItems,
	}
}

// Creates a new booking and assigns the current user as the owner.
func (s *bookingService) CreateBooking(
	ctx context.Context,
	req booking.CreateBookingRequest,
) booking.CreateBookingResponse {
	orgID := booking.OrganizationIDFromContext(ctx)
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return booking.CreateBookingResponse{
			Err: fmt.Errorf("failed to start transaction: %w", err),
		}
	}

	// Get resource that booking will be created for.
	r, err := tx.Resource.
		Query().
		Where(
			resource.ID(req.ResourceID),
			resource.OrganizationId(orgID),
		).
		First(ctx)
	var nfe *NotFoundError
	if errors.As(err, &nfe) {
		return booking.CreateBookingResponse{
			Err: booking.WrapNotFoundError("resource"),
		}
	}
	if err != nil {
		return booking.CreateBookingResponse{
			Err: fmt.Errorf("failed to get resource: %w", err),
		}
	}

	// If quantity available is nil then there is no limit to the number of
	// bookings that can be made for a resource.
	if r.QuantityAvailable != nil {
		// Ensure that quantity of bookings available for resource has not been
		// exceeded.
		if count, err := countOverlappingBookings(
			ctx,
			tx,
			req.ResourceID,
			req.StartTime,
			req.EndTime,
		); err != nil {
			return booking.CreateBookingResponse{
				Err: fmt.Errorf("failed to count overlapping bookings: %w", err),
			}
		} else if count >= *r.QuantityAvailable {
			return booking.CreateBookingResponse{
				Err: booking.Error{
					Code:   booking.ECONFLICT,
					Detail: "Maximum bookings for specified resource reached.",
					Title:  "Resource unavailable",
				},
			}
		}
	}

	b, err := createBooking(ctx, tx, req)
	if err != nil {
		return booking.CreateBookingResponse{
			Err: fmt.Errorf("failed to create booking: %w", err),
		}
	}

	if err := b.attachEdges(ctx); err != nil {
		return booking.CreateBookingResponse{
			Err: err,
		}
	}

	return booking.CreateBookingResponse{
		Booking: b.toModel(),
		Err:     tx.Commit(),
	}
}

func countOverlappingBookings(ctx context.Context, tx *Tx, resourceID int, startTime time.Time, endTime time.Time) (int, error) {
	return tx.Booking.
		Query().
		Where(
			entbooking.ResourceId(resourceID),
			entbooking.Or(
				// New booking begins during an existing booking.
				entbooking.And(entbooking.StartTimeLTE(startTime), entbooking.EndTimeGTE(startTime)),
				// New booking ends during an exsting booking.
				entbooking.And(entbooking.StartTimeLTE(endTime), entbooking.EndTimeGTE(endTime)),
				// New booking is entirely during an existing booking.
				entbooking.And(entbooking.StartTimeLTE(startTime), entbooking.EndTimeGTE(endTime)),
				// Existing booking is entirely during new booking.
				entbooking.And(entbooking.StartTimeGTE(startTime), entbooking.EndTimeLTE(endTime)),
			),
		).
		Count(ctx)
}

func createBooking(ctx context.Context, tx *Tx, req booking.CreateBookingRequest) (*Booking, error) {
	var m []*BookingMetadatum
	for k, v := range req.Metadata {
		m = append(m, &BookingMetadatum{
			Key:   k,
			Value: v,
		})
	}

	return tx.Booking.
		Create().
		SetResourceID(req.ResourceID).
		SetStatus(req.Status).
		SetStartTime(req.StartTime).
		SetEndTime(req.EndTime).
		AddMetadata(m...).
		Save(ctx)
}

func (b *Booking) attachEdges(ctx context.Context) (err error) {
	b.Edges.Metadata, err = b.QueryMetadata().All(ctx)
	if err != nil {
		return err
	}
	b.Edges.Resource, err = b.QueryResource().First(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Updates an existing booking by ID. Only the booking owner can update a
// booking. Returns the new booking state even if there was an error during
// update.
//
// Returns ENOTFOUND if the booking does not exist or the user does not have
// permission to update it.
func (s *bookingService) UpdateBooking(
	ctx context.Context,
	req booking.UpdateBookingRequest,
) booking.UpdateBookingResponse {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return booking.UpdateBookingResponse{
			Err: fmt.Errorf("failed to start transaction: %w", err),
		}
	}

	b, err := findBookingByID(ctx, tx, req.ID)
	var nfe *NotFoundError
	if errors.As(err, &nfe) {
		return booking.UpdateBookingResponse{
			Err: booking.WrapNotFoundError("booking"),
		}
	}
	if err != nil {
		return booking.UpdateBookingResponse{
			Err: fmt.Errorf("failed to find booking by id: %w", err),
		}
	}
}

func updateBooking(ctx context.Context, tx *Tx, req booking.UpdateBookingRequest) error {
	orgID := booking.OrganizationIDFromContext(ctx)
	tx.Booking.
		Update().
		Where(
			entbooking.ID(req.ID),
			entbooking.HasResourceWith(resource.OrganizationId(orgID)),
		).
		Save(ctx)
}

// Permanently removes a booking by ID. Only the booking owner may delete a
// booking. Returns ENOTFOUND if the booking does not exist or the user does
// not have permission to delete it.
func (s *bookingService) DeleteBooking(
	ctx context.Context,
	req booking.DeleteBookingRequest,
) booking.DeleteBookingResponse {
	panic("not implemented") // TODO: Implement
}

func (b *Booking) toModel() *booking.Booking {
	result := &booking.Booking{
		ID:         b.ID,
		ResourceID: b.ResourceId,
		Status:     b.Status,
		StartTime:  b.StartTime,
		EndTime:    b.EndTime,
		CreatedAt:  b.CreatedAt,
		UpdatedAt:  b.UpdatedAt,
	}

	if b.Edges.Resource != nil {
		result.Resource = b.Edges.Resource.toModel()
	}

	if b.Edges.Metadata != nil {
		result.Metadata = BookingMetadata(b.Edges.Metadata).toMap()
	}

	return result
}

func (b Bookings) toModels() []*booking.Booking {
	var bookings []*booking.Booking
	for _, v := range b {
		bookings = append(bookings, v.toModel())
	}
	return bookings
}

func (m BookingMetadata) toMap() map[string]string {
	metadata := make(map[string]string)
	for i := range m {
		metadata[m[i].Key] = m[i].Value
	}
	return metadata
}
