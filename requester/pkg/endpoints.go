package requester

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

// request
type Request struct {
	Coordinates []Point
}

// TripMetrics response
type tripMetricsResponse struct {
	Distance int    `json:"distance"`
	Duration int    `json:"duration"`
	Err      string `json:"err,omitempty"`
}

func MakeTripMetricsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(Request)
		t, d, err := svc.TripMetrics(req.Coordinates)
		return tripMetricsResponse{Distance: d, Duration: t, Err: ""}, err
	}
}
