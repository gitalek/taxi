package calc

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
)

// request api v1
type Request struct {
	Coordinates [][]float64 `json:"coordinates"`
	Strategy    int         `json:"strategy"`
	Rate        string      `json:"rate"`
}

// request api v2
type RequestV2 struct {
	Coordinates []Point `json:"coordinates"`
	Strategy    int     `json:"strategy"`
	Rate        string  `json:"rate"`
}

type Point struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// response
type Response struct {
	Price float64 `json:"price"`
	Err   string  `json:"err,omitempty"`
}

func MakeCalculatePriceEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(RequestV2)
		if !ok {
			err = errors.New("MakeCalculatePriceEndpoint: error while Request type casting")
			return Response{Err: err.Error()}, nil
		}
		price, err := svc.Price(ctx, req.Coordinates, req.Strategy, req.Rate)
		if err != nil {
			return Response{Err: err.Error()}, nil
		}
		return Response{Price: price}, nil
	}
}
