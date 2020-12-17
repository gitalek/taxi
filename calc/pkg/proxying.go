package calc

import (
	"errors"
	"github.com/go-kit/kit/endpoint"
)

type proxymw struct {
	next Service
	tripMetrics endpoint.Endpoint
}

// todo ввести err
func (p proxymw) TripMetrics(c []Point) (int, int, error) {
	response, err := p.tripMetrics(nil, Request{Coordinates: c})
	if err != nil {
		return 0, 0, err
	}
	resp := response.(tripMetricsResponse)
	if resp.Err != "" {
		return resp.Distance, resp.Duration, errors.New(resp.Err)
	}

	return resp.Distance, resp.Duration, nil
}
