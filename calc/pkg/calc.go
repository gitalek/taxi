// service business logic
package calc

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

// service as an interface
type Service interface {
	Price(context.Context, []Point) (int, error)
	CalculatePrice(context.Context, int, int) int
	TripMetrics(context.Context, BusinessRequest) (int, int, error)
}

const (
	taxiService = 50
	minPrice    = 150
	minuteRate  = 10
	kmRate      = 20
)

const apiUrl = "http://localhost:9091/tripmetrics"

type BusinessRequest struct {
	Coordinates []Point
}

func (r BusinessRequest) ORSRequest() ORSRequest {
	coordinates := make([][]float64, 0, len(r.Coordinates))
	for _, point := range r.Coordinates {
		coordinates = append(coordinates, []float64{point.Lat, point.Lon})
	}
	return ORSRequest{coordinates}
}

type ORSRequest struct {
	Coordinates [][]float64 `json:"coordinates"`
}

// TripMetrics response
type MetricsResponse struct {
	Distance int    `json:"distance"`
	Duration int    `json:"duration"`
	Err      string `json:"err,omitempty"`
}

// implementation of the interface
type CalcService struct{}

// check interface realization
var _ Service = &CalcService{}

func (s CalcService) Price(ctx context.Context, c []Point) (int, error) {
	request := BusinessRequest{c}
	t, d, err := s.TripMetrics(ctx, request)
	if err != nil {
		return 0, err
	}
	price := s.CalculatePrice(ctx, t, d)
	return price, err
}

// CalculatePrice calculate a price of the trip in rubles (int);
// params: t - number of minutes (int), dist - number of kilometers (int)
func (s *CalcService) CalculatePrice(ctx context.Context, t int, dist int) int {
	actualPrice := taxiService + t*minuteRate + dist*kmRate
	if minPrice >= actualPrice {
		return minPrice
	}
	return actualPrice
}

// tripMetrics is a temporary stub method until API2 realization
func (*CalcService) TripMetrics(ctx context.Context, request BusinessRequest) (int, int, error) {
	client := &http.Client{}
	body, err := json.Marshal(request.ORSRequest())
	if err != nil {
		log.Println("Errored while marshalling")
		return 0, 0, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", apiUrl, bytes.NewBuffer(body))
	if err != nil {
		log.Println("Errored when create request to the server")
		return 0, 0, err
	}
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Errored when sending request to the server")
		return 0, 0, err
	}
	defer resp.Body.Close()

	var metrics MetricsResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&metrics)
	if err != nil {
		log.Printf("Errored while decoding: %#v\n", err)
	}
	duration := metrics.Duration
	dist := metrics.Distance
	return duration, dist, nil
}

// todo ?
// ServiceMiddleware is a chainable behaviour modifier for Service
type ServiceMiddleware func(Service) Service
