package requester

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

const authKey = "5b3ce3597851110001cf624808fb2d8f3a0048a6b091b55a9da65187"
const apiUrl = "https://api.openrouteservice.org/v2/directions/driving-car/geojson"

// service as an interface
type Service interface {
	TripMetrics(context.Context, []Point) (int, int, error)
}

// implementation of the interface
type RequesterService struct{}
// check interface realization
var _ Service = &RequesterService{}

// tripMetrics is a temporary stub method until API2 realization
func (*RequesterService) TripMetrics(ctx context.Context, c []Point) (int, int, error) {
	request := BusinessMessage{c}
	client := &http.Client{}
	body, err := json.Marshal(request.ORSRequest())
	if err != nil {
		log.Println("Errored while marshalling")
		return 0, 0, err
	}
	req, _ := http.NewRequestWithContext(ctx, "POST", apiUrl, bytes.NewBuffer(body))
	req.Header.Add("Accept", "application/json, application/geo+json, application/gpx+xml, img/png; charset=utf-8")
	req.Header.Add("Authorization", authKey)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Errored when sending request to the server")
		return 0, 0, err
	}
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
	duration:= metrics.Features[0].Properties.Summary.Duration
	dist := metrics.Features[0].Properties.Summary.Distance
	return int(duration), int(dist), nil
}
