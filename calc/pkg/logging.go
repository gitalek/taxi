package calc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
)

type Middleware func(endpoint.Endpoint) endpoint.Endpoint

func LoggingMiddleware(logger *zap.SugaredLogger) Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			logger.Infow("logging middleware called", "params: request", request)
			return next(ctx, request)
		}
	}
}

type AppLoggingMiddleware struct {
	Logger *zap.SugaredLogger
	Next   Service
}

var _ Service = &AppLoggingMiddleware{}

func (mv AppLoggingMiddleware) Price(ctx context.Context, c []Point) (int, error) {
	mv.Logger.Infow(
		"",
		"method", "Price",
		"params: c", c,
	)
	return mv.Next.Price(ctx, c)
}

//todo useless logging
//func (mv AppLoggingMiddleware) tripMetrics(ctx context.Context, message BusinessMessage) (int, int, error) {
//	logger := log.With(mv.Logger, "method", "TripMetrics")
//	err := logger.Log("in: message", message)
//	if err != nil {
//		l.Printf("calc: logic: Price: error while logging: %#v\n", err)
//	}
//	return mv.Next.tripMetrics(ctx, message)
//}

//todo useless logging
//func (mv AppLoggingMiddleware) calculatePrice(ctx context.Context, t int, dist int) int {
//	logger := log.With(mv.Logger, "method", "CalculatePrice")
//	err := logger.Log("in: time", t, "in: distance", dist)
//	if err != nil {
//		l.Printf("calc: logic: Price: error while logging: %#v\n", err)
//	}
//	return mv.Next.calculatePrice(ctx, t, dist)
//}
