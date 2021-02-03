package _map

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/gitalek/taxi/requester/pkg/types"
	"log"
	"net/http"
)

type ORSRequest struct {
	Coordinates [][]float64 `json:"coordinates"`
	Language    string      `json:"language"`
	Units       string      `json:"units"`
}

type ORSResponse struct {
	Features []struct {
		Properties struct {
			Summary struct {
				Distance float64 `json:"distance"`
				Duration float64 `json:"duration"`
			} `json:"summary"`
		} `json:"properties"`
	} `json:"features"`
}

func ORSMetrics(ctx context.Context, c []types.Point, key string, url string) (float64, float64, error) {
	req, err := prepareORSRequest(ctx, c, key, url)
	if err != nil {
		return 0, 0, err
	}

	//todo global client?
	client := &http.Client{}
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
		log.Printf("ErrNoStructureProperty: %#v\n", metrics)
		return 0, 0, errors.New("no data error") //todo: ввести кастомный тип ошибки?
	}
	//todo: проверить наличие свойств по цепочке ".Properties.Summary.Duration"
	duration := metrics.Features[0].Properties.Summary.Duration
	dist := metrics.Features[0].Properties.Summary.Distance
	log.Printf("duration -> %#v, dist -> %#v\n", duration, dist)
	return duration, dist, nil
}

func prepareORSRequest(ctx context.Context, points []types.Point, key string, url string) (req *http.Request, err error) {
	coordinates := make([][]float64, 0, len(points))
	for _, point := range points {
		// Start coordinate of the route in longitude,latitude format. - openroute
		//coordinates = append(coordinates, []float64{point.Lat, point.Lon})
		coordinates = append(coordinates, []float64{point.Lon, point.Lat})
	}
	r := ORSRequest{Coordinates: coordinates, Language: "ru", Units: "m"}
	body, err := json.Marshal(r)
	if err != nil {
		log.Printf("Errored while marshalling: %#v\n", err)
		return
	}
	req, err = http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Errored when create request to the server: %#v\n", err)
		return
	}
	req.Header.Add("Accept", "application/json, application/geo+json, application/gpx+xml, img/png; charset=utf-8")
	// todo check key existence
	req.Header.Add("Authorization", key)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	return req, nil
}
