package requester

import (
	"context"
	"errors"
	"github.com/gitalek/taxi/requester/pkg/types"
	"github.com/go-kit/kit/endpoint"
)

type Request struct {
	Coordinates []types.Point
	Strategy    int
}

// TripMetrics response
type tripMetricsResponse struct {
	Distance float64 `json:"distance"`
	Duration float64 `json:"duration"`
	Err      string  `json:"err,omitempty"`
}

func MakeTripMetricsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(Request)
		if !ok {
			err = errors.New("MakeTripMetricsEndpoint: error while Request type casting")
			return tripMetricsResponse{Err: err.Error()}, nil
		}
		t, d, err := svc.TripMetrics(ctx, req.Coordinates, req.Strategy)
		if err != nil {
			return tripMetricsResponse{Err: err.Error()}, nil
		}
		return tripMetricsResponse{Distance: d, Duration: t, Err: ""}, nil
	}
}
