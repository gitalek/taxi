package calc

import (
	"context"
	"encoding/json"
	"net/http"
)

func DecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	message := transformToMessage(req)
	return message, nil
}

func transformToMessage(req Request) RequestV2 {
	points := make([]Point, 0, len(req.Coordinates))
	for _, p := range req.Coordinates {
		point := Point{Lat: p[0], Lon: p[1]}
		points = append(points, point)
	}
	return RequestV2{Coordinates: points}
}

func DecodeRequestV2(_ context.Context, r *http.Request) (interface{}, error) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
