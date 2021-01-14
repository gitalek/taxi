package requester

import (
	"errors"
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
		req, ok := request.(Request)
		if !ok {
			//todo: как обработать ошибку?
			err = errors.New("MakeTripMetricsEndpoint: error while Request type casting")
			return tripMetricsResponse{Err: err.Error()}, err
		}
		log.Printf("TripMetricsEndpoint: request: %#v\n", req)
		t, d, err := svc.TripMetrics(req.Coordinates)
		if err != nil {
			return tripMetricsResponse{Err: err.Error()}, err
		}
		return tripMetricsResponse{Distance: d, Duration: t, Err: ""}, nil
	}
}
