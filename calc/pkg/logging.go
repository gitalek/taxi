package calc

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
				l.Printf("calc: endpoint: calculatePrice: error while logging: %#v\n", err)
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

func (mv AppLoggingMiddleware) Price(ctx context.Context, c []Point) (int, error) {
	logger := log.With(mv.Logger, "method", "Price")
	err := logger.Log("params: c", c, "ts", log.DefaultTimestampUTC)
	if err != nil {
		l.Printf("calc: logic: Price: error while logging: %#v\n", err)
	}
	return mv.Next.Price(ctx, c)
}

//todo useless logging
func (mv AppLoggingMiddleware) tripMetrics(ctx context.Context, message BusinessMessage) (int, int, error) {
	logger := log.With(mv.Logger, "method", "TripMetrics")
	err := logger.Log("in: message", message)
	if err != nil {
		l.Printf("calc: logic: Price: error while logging: %#v\n", err)
	}
	return mv.Next.tripMetrics(ctx, message)
}

//todo useless logging
func (mv AppLoggingMiddleware) calculatePrice(ctx context.Context, t int, dist int) int {
	logger := log.With(mv.Logger, "method", "CalculatePrice")
	err := logger.Log("in: time", t, "in: distance", dist)
	if err != nil {
		l.Printf("calc: logic: Price: error while logging: %#v\n", err)
	}
	return mv.Next.calculatePrice(ctx, t, dist)
}
