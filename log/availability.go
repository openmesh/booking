package log

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/openmesh/booking"
)

func AvailabilityLoggingMiddleware(logger log.Logger) booking.AvailabilityServiceMiddleware {
	return func(next booking.AvailabilityService) booking.AvailabilityService {
		return availabilityLoggingMiddleware{logger, next}
	}
}

type availabilityLoggingMiddleware struct {
	logger log.Logger
	booking.AvailabilityService
}

func (mw availabilityLoggingMiddleware) FindAvailabilities(ctx context.Context, filter booking.AvailabilityFilter) (availabilities []*booking.Availability, totalItems int, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "find_availabilities",
			"filter", filter,
			"availabilities", availabilities,
			"total_items", totalItems,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	availabilities, totalItems, err = mw.AvailabilityService.FindAvailabilties(ctx, filter)
	return
}
