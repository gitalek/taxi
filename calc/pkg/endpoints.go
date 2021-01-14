package calc

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
)

type Point struct {
	Lat float64
	Lon float64
}

// request
type Request struct {
	Coordinates [][]float64
}

// response
type Response struct {
	Price int    `json:"price"`
	Err   string `json:"err,omitempty"`
}

func MakeCalculatePriceEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(Request)
		if !ok {
			//todo: как обработать ошибку?
			return nil, errors.New("MakeCalculatePriceEndpoint: error while Request type casting")
		}
		t, d, err := svc.TripMetrics(req.Coordinates)
		if err != nil {
			return nil, err
		}
		price := svc.CalculatePrice(t, d)
		return Response{price, ""}, nil
	}
}
