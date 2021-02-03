package _map

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gitalek/taxi/requester/pkg/types"
	"log"
	"net/http"
)

type BMResponse struct {
	ResourceSets []struct {
		Resources []struct {
			TravelDistance        float64 `json:"travelDistance"`
			TravelDuration        float64     `json:"travelDuration"`
			TravelDurationTraffic float64 `json:"travelDurationTraffic"`
		} `json:"resources"`
	} `json:"resourceSets"`
}

func BingMapsMetrics(ctx context.Context, c []types.Point, key string, url string) (float64, float64, error) {
	req, err := prepareBingMapsRequest(ctx, c, key, url)
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
	duration := metrics.ResourceSets[0].Resources[0].TravelDurationTraffic
	//duration := metrics.ResourceSets[0].Resources[0].TravelDuration
	dist := metrics.ResourceSets[0].Resources[0].TravelDistance
	// convert from km to meters
	dist = dist * 1000
	log.Printf("duration -> %#v, dist -> %#v\n", duration, dist)
	return duration, dist, nil
}

func prepareBingMapsRequest(ctx context.Context, points []types.Point, key string, url string) (req *http.Request, err error) {
	// This example gets address information for a specified latitude and longitude and requests the results in XML format. - bing
	call := fmt.Sprintf(
		"%s?wp.0=%f,%f&wp.1=%f,%f&key=%s&du=km",
		url, points[0].Lat, points[0].Lon, points[1].Lat, points[1].Lon, key,
	)
	req, err = http.NewRequestWithContext(ctx, "GET", call, nil)
	if err != nil {
		log.Printf("Errored when create request to the server: %#v\n", err)
		return
	}
	return req, nil
}
