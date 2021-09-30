package event

import (
	"context"

	"github.com/openmesh/booking"
)

func BookingMiddleware(eventService booking.EventService) booking.BookingServiceMiddleware {
	return func(next booking.BookingService) booking.BookingService {
		return bookingEventMiddleware{eventService, next}
	}
}

type bookingEventMiddleware struct {
	booking.EventService
	booking.BookingService
}

// Retrieves a single booking by ID along with the associated resource and
// metadata. Returns ENOTFOUND if booking does not exist or user does not have
// permission to view it.
func (mw bookingEventMiddleware) FindBookingByID(ctx context.Context, req booking.FindBookingByIDRequest) booking.FindBookingByIDResponse {
	return mw.BookingService.FindBookingByID(ctx, req)
}

// Retreives a list of bookings based on a filter. Only returns bookings that
// are accessible to the user. Also returns a count of total matching bookings
// which may be different from the number of returned bookings if the "Limit"
// field is set.
func (mw bookingEventMiddleware) FindBookings(ctx context.Context, req booking.FindBookingsRequest) booking.FindBookingsResponse {
	return mw.BookingService.FindBookings(ctx, req)
}

// Creates a new booking and assigns the current user as the owner.
func (mw bookingEventMiddleware) CreateBooking(ctx context.Context, req booking.CreateBookingRequest) (res booking.CreateBookingResponse) {
	defer func() {
		ev := booking.Event{
			Type:    booking.EventTypeBookingCreated,
			Payload: booking.BookingCreatedPayload{Booking: res.Booking},
		}
		userID := booking.UserIDFromContext(ctx)
		mw.EventService.PublishEvent(userID, ev)
	}()
	res = mw.BookingService.CreateBooking(ctx, req)
	return
}

// Updates an existing booking by ID. Only the booking owner can update a
// booking. Returns the new booking state even if there was an error during
// update.
//
// Returns ENOTFOUND if the booking does not exist or the user does not have
// permission to update it.
func (mw bookingEventMiddleware) UpdateBooking(ctx context.Context, req booking.UpdateBookingRequest) (res booking.UpdateBookingResponse) {
	defer func() {
		ev := booking.Event{
			Type:    booking.EventTypeBookingUpdated,
			Payload: booking.BookingUpdatedPayload{Booking: res.Booking},
		}
		userID := booking.UserIDFromContext(ctx)
		mw.EventService.PublishEvent(userID, ev)
	}()
	res = mw.BookingService.UpdateBooking(ctx, req)
	return
}

// Permanently removes a booking by ID. Only the booking owner may delete a
// booking. Returns ENOTFOUND if the booking does not exist or the user does
// not have permission to delete it.
func (mw bookingEventMiddleware) DeleteBooking(ctx context.Context, req booking.DeleteBookingRequest) (res booking.DeleteBookingResponse) {
	defer func() {
		ev := booking.Event{
			Type:    booking.EventTypeBookingDeleted,
			Payload: booking.BookingDeletedPayload{ID: req.ID},
		}
		userID := booking.UserIDFromContext(ctx)
		mw.EventService.PublishEvent(userID, ev)
	}()
	res = mw.BookingService.DeleteBooking(ctx, req)
	return
}
