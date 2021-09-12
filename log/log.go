package log
import (
	"github.com/go-kit/kit/log"
	"github.com/openmesh/booking"
)

func LoggingMiddleware(logger log.Logger) booking.HandlerMiddleware {
	return func(next booking.Handler) booking.Handler {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	booking.Handler
}