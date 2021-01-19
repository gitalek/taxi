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
}

const (
	taxiService = 50
	minPrice    = 150
	minuteRate  = 10
	kmRate      = 20
)

const apiUrl = "http://localhost:9091/tripmetrics"

// implementation of the interface
type CalcService struct{}

// check interface realization
var _ Service = &CalcService{}

func (s CalcService) Price(ctx context.Context, c []Point) (int, error) {
	message := BusinessMessage{c}
	t, d, err := s.tripMetrics(ctx, message)
	if err != nil {
		return 0, err
	}
	price := s.calculatePrice(ctx, t, d)
	return price, err
}

// CalculatePrice calculate a price of the trip in rubles (int);
// params: t - number of minutes (int), dist - number of meters (int)
func (s *CalcService) calculatePrice(ctx context.Context, t int, dist int) int {
	// todo check number types
	actualPrice := taxiService + t*minuteRate + (dist * kmRate / 1000)
	if minPrice >= actualPrice {
		return minPrice
	}
	return actualPrice
}

// tripMetrics is a temporary stub method until API2 realization
func (*CalcService) tripMetrics(ctx context.Context, message BusinessMessage) (int, int, error) {
	client := &http.Client{}
	body, err := json.Marshal(message.Request())
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

	metrics := message.Response()
	err = json.NewDecoder(resp.Body).Decode(&metrics)
	if err != nil {
		log.Printf("Errored while decoding: %#v\n", err)
		return 0, 0, err
	}
	// todo check service response struct error
	// todo multiple points case
	duration := metrics.Duration
	dist := metrics.Distance
	return duration, dist, nil
}

// todo ?
// ServiceMiddleware is a chainable behaviour modifier for Service
type ServiceMiddleware func(Service) Service
