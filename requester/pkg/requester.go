package requester

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/gitalek/taxi/requester/pkg/map"
	"log"
	"net/http"
)

// service as an interface
type Service interface {
	TripMetrics(context.Context, []Point) (float64, float64, error)
}

type ServiceConfig struct {
	ApiUrl string
	ORSKey string
}

// implementation of the interface
type RequesterService struct {
	Config ServiceConfig
}

// check interface realization
var _ Service = &RequesterService{}

func (s *RequesterService) TripMetrics(ctx context.Context, c []Point) (float64, float64, error) {
	return _map.BingMapsMetrics(ctx, c, s.Config.ORSKey, s.Config.ApiUrl)
}

func (s *RequesterService) _TripMetrics(ctx context.Context, c []Point) (int, int, error) {
	request := BusinessMessage{c}
	//todo global client?
	client := &http.Client{}
	body, err := json.Marshal(request.ORSRequest())
	if err != nil {
		log.Printf("Errored while marshalling: %#v\n", err)
		return 0, 0, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", s.Config.ApiUrl, bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Errored when create request to the server: %#v\n", err)
		return 0, 0, err
	}
	req.Header.Add("Accept", "application/json, application/geo+json, application/gpx+xml, img/png; charset=utf-8")
	//todo check key existence
	req.Header.Add("Authorization", s.Config.ORSKey)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Errored when sending request to the server: %#v\n", err)
		return 0, 0, err
	}
	//todo handle error
	defer resp.Body.Close()

	var metrics ORSResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&metrics)
	if err != nil {
		log.Printf("TripMetrics: errored while decoding: %#v\n", err)
		return 0, 0, err
	}
	if len(metrics.Features) == 0 {
		return 0, 0, errors.New("no data error") //todo: ввести кастомный тип ошибки?
	}
	//todo: проверить наличие свойств по цепочке ".Properties.Summary.Duration"
	duration := metrics.Features[0].Properties.Summary.Duration
	dist := metrics.Features[0].Properties.Summary.Distance
	return int(duration), int(dist), nil
}

func (s *RequesterService) ORSMetrics(ctx context.Context, c []Point) (int, int, error) {
	request := BusinessMessage{c}
	//todo global client?
	client := &http.Client{}
	body, err := json.Marshal(request.ORSRequest())
	if err != nil {
		log.Printf("Errored while marshalling: %#v\n", err)
		return 0, 0, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", s.Config.ApiUrl, bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Errored when create request to the server: %#v\n", err)
		return 0, 0, err
	}
	req.Header.Add("Accept", "application/json, application/geo+json, application/gpx+xml, img/png; charset=utf-8")
	// todo check key existence
	req.Header.Add("Authorization", s.Config.ORSKey)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Errored when sending request to the server: %#v\n", err)
		return 0, 0, err
	}
	//todo handle error
	defer resp.Body.Close()

	var metrics ORSResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&metrics)
	if err != nil {
		log.Printf("TripMetrics: errored while decoding: %#v\n", err)
		return 0, 0, err
	}
	if len(metrics.Features) == 0 {
		return 0, 0, errors.New("no data error") //todo: ввести кастомный тип ошибки?
	}
	//todo: проверить наличие свойств по цепочке ".Properties.Summary.Duration"
	duration := metrics.Features[0].Properties.Summary.Duration
	dist := metrics.Features[0].Properties.Summary.Distance
	return int(duration), int(dist), nil
}

func (s *RequesterService) BingMapsMetrics(ctx context.Context, c []Point) (int, int, error) {
	request := BusinessMessage{c}
	//todo global client?
	client := &http.Client{}
	body, err := json.Marshal(request.ORSRequest())
	if err != nil {
		log.Printf("Errored while marshalling: %#v\n", err)
		return 0, 0, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", s.Config.ApiUrl, bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Errored when create request to the server: %#v\n", err)
		return 0, 0, err
	}
	req.Header.Add("Accept", "application/json, application/geo+json, application/gpx+xml, img/png; charset=utf-8")
	//todo check key existence
	req.Header.Add("Authorization", s.Config.ORSKey)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Errored when sending request to the server: %#v\n", err)
		return 0, 0, err
	}
	//todo handle error
	defer resp.Body.Close()

	var metrics ORSResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&metrics)
	if err != nil {
		log.Printf("TripMetrics: errored while decoding: %#v\n", err)
		return 0, 0, err
	}
	if len(metrics.Features) == 0 {
		return 0, 0, errors.New("no data error") //todo: ввести кастомный тип ошибки?
	}
	//todo: проверить наличие свойств по цепочке ".Properties.Summary.Duration"
	duration := metrics.Features[0].Properties.Summary.Duration
	dist := metrics.Features[0].Properties.Summary.Distance
	return int(duration), int(dist), nil
}
