package calc

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
)

// request api v1
type Request struct {
	Coordinates [][]float64 `json:"coordinates"`
}

// request api v2
type RequestV2 struct {
	Coordinates []Point `json:"coordinates"`
}

type Point struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// response
type Response struct {
	Price int    `json:"price"`
	Err   string `json:"err,omitempty"`
}

func MakeCalculatePriceEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(RequestV2)
		if !ok {
			err = errors.New("MakeCalculatePriceEndpoint: error while Request type casting")
			return Response{Err: err.Error()}, err
		}
		price, err := svc.Price(ctx, req.Coordinates)
		if err != nil {
			return Response{Err: err.Error()}, err
		}
		return Response{Price: price}, nil
	}
}
