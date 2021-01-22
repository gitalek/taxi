package _map

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	requester "github.com/gitalek/taxi/requester/pkg"
	"log"
	"net/http"
)

type BMResponse struct {
	ResourceSets []struct {
		Resources []struct {
			TravelDistance float64 `json:"travelDistance"`
			TravelDuration float64 `json:"travelDuration"`
		} `json:"resources"`
	} `json:"resourceSets"`
}

func BingMapsMetrics(ctx context.Context, c []requester.Point, key string, url string) (float64, float64, error) {
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

	var metrics BMResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&metrics)
	if err != nil {
		log.Printf("TripMetrics: errored while decoding: %#v\n", err)
		return 0, 0, err
	}
	//todo: проверить наличие свойств по цепочке
	duration := metrics.ResourceSets[0].Resources[0].TravelDuration
	dist := metrics.ResourceSets[0].Resources[0].TravelDistance
	return duration, dist, nil
}

func prepareBingMapsRequest(ctx context.Context, points []requester.Point, key string, url string) (req *http.Request, err error) {
	call := fmt.Sprintf(
		"%s?wp.0=%f,%f&wp.1=%f,%f&key=%s'",
		url, points[0].Lat, points[0].Lon, points[1].Lat, points[1].Lon, key,
	)
	req, err = http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer([]byte(call)))
	if err != nil {
		log.Printf("Errored when create request to the server: %#v\n", err)
		return
	}
	return req, nil
}
