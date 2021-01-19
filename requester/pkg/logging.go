package requester

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	l "log"
)

type Middleware func(endpoint.Endpoint) endpoint.Endpoint

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			err = logger.Log("params: request", request)
			if err != nil {
				l.Printf("calc: endpoint: TripMetrics: error while logging: %#v\n", err)
			}
			return next(ctx, request)
		}
	}
}

type AppLoggingMiddleware struct {
	Logger log.Logger
	Next   Service
}
var _ Service = &AppLoggingMiddleware{}

func (mv AppLoggingMiddleware) TripMetrics(ctx context.Context, c []Point) (int, int, error) {
	logger := log.With(mv.Logger, "method", "TripMetrics")
	err := logger.Log("params: c", c, "ts", log.DefaultTimestampUTC)
	if err != nil {
		l.Printf("calc: logic: TripMetrics: error while logging: %#v\n", err)
	}
	return mv.Next.TripMetrics(ctx, c)
}
