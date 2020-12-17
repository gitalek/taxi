package calc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"time"
)

type Middleware func(endpoint.Endpoint) endpoint.Endpoint

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			logger.Log("msg", "called endpoint")
			return next(ctx, request)
		}
	}
}

type AppLoggingMiddleware struct {
	Logger log.Logger
	Next Service
}

func (mv AppLoggingMiddleware) TripMetrics(c []Point) (int, int) {
	return mv.Next.TripMetrics(c)
}

func (mv AppLoggingMiddleware) CalculatePrice(t int, dist int) int {
	defer func(begin time.Time) {
		mv.Logger.Log("time: ", t, ", distance: ", dist, ", last: ", time.Since(begin).Seconds())
	}(time.Now())
	return mv.Next.CalculatePrice(t, dist)
}
