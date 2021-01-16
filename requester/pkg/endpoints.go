package requester

import (
	"errors"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

type Request struct {
	Coordinates []Point
}

type Point struct {
	Lat float64
	Lon float64
}

// TripMetrics response
type tripMetricsResponse struct {
	Distance int    `json:"distance"`
	Duration int    `json:"duration"`
	Err      string `json:"err,omitempty"`
}

func MakeTripMetricsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(Request)
		if !ok {
			err = errors.New("MakeTripMetricsEndpoint: error while Request type casting")
			return tripMetricsResponse{Err: err.Error()}, err
		}
		t, d, err := svc.TripMetrics(ctx, req.Coordinates)
		if err != nil {
			return tripMetricsResponse{Err: err.Error()}, err
		}
		return tripMetricsResponse{Distance: d, Duration: t, Err: ""}, nil
	}
}
