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

func (mv AppLoggingMiddleware) Price(ctx context.Context, c []Point, strategy int, rate string) (float64, error) {
	mv.Logger.Infow(
		"",
		"method", "Price",
		"params: c", c,
		"strategy", strategy,
		"rate", rate,
	)
	return mv.Next.Price(ctx, c, strategy, rate)
}

