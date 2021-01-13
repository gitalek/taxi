package requester

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
	"log"
)

// request
type Request struct {
	Coordinates [][]float64 `json:"coordinates"`
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
		log.Printf("TripMetricsEndpoint: request: %#v\n", req)
		t, d, err := svc.TripMetrics(req.Coordinates)
		return tripMetricsResponse{Distance: d, Duration: t, Err: ""}, err
	}
}
