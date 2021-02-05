package requester

import (
	"context"
	"github.com/gitalek/taxi/requester/pkg/types"
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

func (mv AppLoggingMiddleware) TripMetrics(ctx context.Context, c []types.Point, strategy int) (float64, float64, error) {
	mv.Logger.Infow(
		"",
		"method", "TripMetrics",
		"params: c", c,
	)
	return mv.Next.TripMetrics(ctx, c, strategy)
}
