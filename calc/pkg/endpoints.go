package calc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Point struct {
	Lat float64
	Lon float64
}

// request
type Request struct {
	Coordinates []Point
}

// response
type Response struct {
	Price int   `json:"price"`
	Err   string `json:"err,omitempty"`
}

func MakeCalculatePriceEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(Request)
		t, d := svc.TripMetrics(req.Coordinates)
		price := svc.CalculatePrice(t, d)
		return Response{price, ""}, nil
	}
}

