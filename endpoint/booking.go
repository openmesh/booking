package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/openmesh/booking"
)

// BookingEndpoints collects all the endpoints that compose a booking.BookingService.
// It's used as a helper struct, to collect all the endpoints into a single
// parameter.
type BookingEndpoints struct {
	FindBookingByIDEndpoint endpoint.Endpoint
	FindBookingsEndpoint    endpoint.Endpoint
	CreateBookingEndpoint   endpoint.Endpoint
	UpdateBookingEndpoint   endpoint.Endpoint
	DeleteBookingEndpoint   endpoint.Endpoint
}

// MakeBookingEndpoints returns a BookingEndpoints struct where each endpoint
// invokes the corresponding method on the provided service.
func MakeBookingEndpoints(s booking.BookingService) BookingEndpoints {
	return BookingEndpoints{
		FindBookingByIDEndpoint: MakeFindBookingByIDEndpoint(s),
		FindBookingsEndpoint:    MakeFindBookingsEndpoint(s),
		CreateBookingEndpoint:   MakeCreateBookingEndpoint(s),
		UpdateBookingEndpoint:   MakeUpdateBookingEndpoint(s),
		DeleteBookingEndpoint:   MakeDeleteBookingEndpoint(s),
	}
}

// MakeFindBookingByIDEndpoint returns an endpoint via the passed service.
func MakeFindBookingByIDEndpoint(s booking.BookingService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		return s.FindBookingByID(ctx, r.(booking.FindBookingByIDRequest)), nil
	}
}

// MakeFindBookingsEndpoint returns an endpoint via the passed service.
func MakeFindBookingsEndpoint(s booking.BookingService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		return s.FindBookings(ctx, r.(booking.FindBookingsRequest)), nil
	}
}

// MakeCreateBookingEndpoint returns an endpoint via the passed service.
func MakeCreateBookingEndpoint(s booking.BookingService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		return s.CreateBooking(ctx, r.(booking.CreateBookingRequest)), nil
	}
}

// MakeUpdateBookingEndpoint returns an endpoint via the passed service.
func MakeUpdateBookingEndpoint(s booking.BookingService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		return s.UpdateBooking(ctx, r.(booking.UpdateBookingRequest)), nil
	}
}

// MakeDeleteBookingEndpoint returns an endpoint via the passed service.
func MakeDeleteBookingEndpoint(s booking.BookingService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		return s.DeleteBooking(ctx, r.(booking.DeleteBookingRequest)), nil
	}
}
