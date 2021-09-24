package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/openmesh/booking"
	entbooking "github.com/openmesh/booking/ent/booking"
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

	b, err := findBookingByID(ctx, tx, req.ID, func(bq *BookingQuery) *BookingQuery {
		return bq.WithResource().WithMetadata()
	})
	if err != nil {
		return booking.FindBookingByIDResponse{
			Err: fmt.Errorf("failed to find booking by id: %w", err),
		}
	}

	return booking.FindBookingByIDResponse{Booking: b.toModel()}
}

func findBookingByID(
	ctx context.Context,
	tx *Tx,
	id int,
	withEdges func(*BookingQuery) *BookingQuery,
) (*Booking, error) {
	q := tx.Booking.
		Query().
		Where(entbooking.ID(id))
	if withEdges != nil {
		withEdges(q)
	}

	b, err := q.First(ctx)
	var nfe *NotFoundError
	if errors.As(err, &nfe) {
		return nil, booking.Errorf(booking.EBOOKINGNOTFOUND, "Could not find booking with ID %d", id)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to find booking: %w", err)
	}

	return b, nil
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

	b, totalItems, err := findBookings(ctx, tx, req, func(bq *BookingQuery) *BookingQuery {
		return bq.WithMetadata().WithResource()
	})
	if err != nil {
		return booking.FindBookingsResponse{
			Err: fmt.Errorf("failed to find bookings: %w", err),
		}
	}

	return booking.FindBookingsResponse{
		Bookings:   Bookings(b).toModels(),
		TotalItems: totalItems,
	}
}

// findBookings is a helper function that retrieves a list of matching bookings.
// Also returns a total matching count which may differ from the number of
// results if a limit is set. The query can be expanded upon using the withEdges
// argument. This is primarily used for eager loading of edges.
func findBookings(
	ctx context.Context,
	tx *Tx,
	req booking.FindBookingsRequest,
	withEdges func(*BookingQuery) *BookingQuery,
) ([]*Booking, int, error) {
	q := tx.Booking.Query()

	if req.ID != nil {
		q.Where(entbooking.ID(*req.ID))
	}
	if req.ResourceID != nil {
		q.Where(entbooking.ResourceId(*req.ResourceID))
	}
	if req.Status != nil {
		q.Where(entbooking.Status(*req.Status))
	}
	if req.StartTimeAfter != nil {
		q.Where(entbooking.StartTimeGTE(*req.StartTimeAfter))
	}
	if req.EndTimeBefore != nil {
		q.Where(entbooking.EndTimeLTE(*req.EndTimeBefore))
	}

	c, err := q.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	q = q.Offset(req.Offset)
	if req.Limit == 0 {
		q = q.Limit(10)
	} else {
		q = q.Limit(req.Limit)
	}

	if withEdges != nil {
		withEdges(q)
	}

	b, err := q.All(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query bookings: %w", err)
	}

	return b, c, nil
}

// Creates a new booking and assigns the current user as the owner.
func (s *bookingService) CreateBooking(
	ctx context.Context,
	req booking.CreateBookingRequest,
) booking.CreateBookingResponse {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return booking.CreateBookingResponse{
			Err: fmt.Errorf("failed to start transaction: %w", err),
		}
	}

	err = checkForBookingTimeConflict(ctx, tx, req.ResourceID, req.StartTime, req.EndTime)
	if err != nil {
		return booking.CreateBookingResponse{
			Err: fmt.Errorf("booking time conflict check failed: %w", err),
		}
	}

	b, err := createBooking(ctx, tx, req, func(b *Booking) (*Booking, error) {
		b.Edges.Metadata, err = b.QueryMetadata().All(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to query metadata")
		}
		b.Edges.Resource, err = b.QueryResource().First(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to query resource")
		}
		return b, nil
	})
	if err != nil {
		return booking.CreateBookingResponse{
			Err: fmt.Errorf("failed to create booking: %w", err),
		}
	}

	err = tx.Commit()
	if err != nil {
		return booking.CreateBookingResponse{
			Err: fmt.Errorf("failed to create booking: %w", err),
		}
	}

	return booking.CreateBookingResponse{
		Booking: b.toModel(),
	}
}

func checkForBookingTimeConflict(
	ctx context.Context,
	tx *Tx,
	rid int,
	st time.Time,
	et time.Time,
	allowedIDs ...int,
) error {
	r, err := findResourceByID(ctx, tx, rid, nil)
	if err != nil {
		return fmt.Errorf("failed to find resource: %w", err)
	}
	// If quantity available is nil then there is no limit to the number of
	// bookings that can made for the specified resource and we can return early.
	if r.QuantityAvailable == nil {
		return nil
	}

	c, err := tx.Booking.
		Query().
		Where(entbooking.IDNotIn(allowedIDs...)).
		Where(
			entbooking.ResourceId(rid),
			entbooking.Or(
				// New booking begins during an existing booking.
				entbooking.And(entbooking.StartTimeLTE(st), entbooking.EndTimeGTE(st)),
				// New booking ends during an exsting booking.
				entbooking.And(entbooking.StartTimeLTE(et), entbooking.EndTimeGTE(et)),
				// New booking is entirely during an existing booking.
				entbooking.And(entbooking.StartTimeLTE(st), entbooking.EndTimeGTE(et)),
				// Existing booking is entirely during new booking.
				entbooking.And(entbooking.StartTimeGTE(st), entbooking.EndTimeLTE(et)),
			),
		).
		Count(ctx)

	if c >= *r.QuantityAvailable {
		return booking.Errorf(
			booking.EBOOKINGCONFLICT,
			"Maximum quantity of bookings allowed for a resource at a single time exceeded",
		)
	}
	return nil
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

func createBooking(
	ctx context.Context,
	tx *Tx,
	req booking.CreateBookingRequest,
	attachEdges func(*Booking) (*Booking, error),
) (*Booking, error) {
	var m []*BookingMetadatum
	for k, v := range req.Metadata {
		m = append(m, &BookingMetadatum{
			Key:   k,
			Value: v,
		})
	}

	b, err := tx.Booking.
		Create().
		SetResourceID(req.ResourceID).
		SetStatus(req.Status).
		SetStartTime(req.StartTime).
		SetEndTime(req.EndTime).
		AddMetadata(m...).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create booking: %w", err)
	}

	b, err = attachEdges(b)
	if err != nil {
		return nil, fmt.Errorf("failed to attach edges: %w", err)
	}
	return b, nil
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

	err = checkForBookingTimeConflict(ctx, tx, req.ResourceID, req.StartTime, req.EndTime)
	if err != nil {
		return booking.UpdateBookingResponse{
			Err: fmt.Errorf("booking time conflict check failed: %w", err),
		}
	}

	b, err := updateBooking(ctx, tx, req, func(b *Booking) (*Booking, error) {
		b.Edges.Resource, err = b.QueryResource().First(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to query resource: %w", err)
		}
		b.Edges.Metadata, err = b.QueryMetadata().All(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to query metadata: %w", err)
		}
		return b, nil
	})

	return booking.UpdateBookingResponse{
		Booking: b.toModel(),
	}
}

func updateBooking(
	ctx context.Context,
	tx *Tx,
	req booking.UpdateBookingRequest,
	attachEdges func(*Booking) (*Booking, error),
) (*Booking, error) {
	b, err := tx.Booking.
		UpdateOneID(req.ID).
		SetStartTime(req.StartTime).
		SetEndTime(req.EndTime).
		SetResourceID(req.ResourceID).
		SetStatus(req.Status).
		Save(ctx)

	var nfe *NotFoundError
	if errors.As(err, &nfe) {
		return nil, booking.Errorf(
			booking.EBOOKINGNOTFOUND,
			"Could not find booking with ID %d",
			req.ID,
		)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to update booking: %w", err)
	}

	b, err = attachEdges(b)
	if err != nil {
		return nil, fmt.Errorf("failed to attach edges: %w", err)
	}

	return b, nil
}

// Permanently removes a booking by ID. Only the booking owner may delete a
// booking. Returns ENOTFOUND if the booking does not exist or the user does
// not have permission to delete it.
func (s *bookingService) DeleteBooking(
	ctx context.Context,
	req booking.DeleteBookingRequest,
) booking.DeleteBookingResponse {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return booking.DeleteBookingResponse{
			Err: fmt.Errorf("failed to start transaction: %w", err),
		}
	}
	err = deleteBooking(ctx, tx, req.ID)
	if err != nil {
		return booking.DeleteBookingResponse{
			Err: fmt.Errorf("failed to delete booking: %w", err),
		}
	}
	return booking.DeleteBookingResponse{}
}

func deleteBooking(ctx context.Context, tx *Tx, id int) error {
	a, err := tx.Booking.
		Delete().
		Where(entbooking.ID(id)).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete booking: %w", err)
	}

	if a == 0 {
		return booking.Errorf(
			booking.EBOOKINGNOTFOUND,
			"Could not find booking with ID %d",
			id,
		)
	}
	return nil
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
