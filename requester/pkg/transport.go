package requester

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func DecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("DecodeRequest: errored while decoding: %#v\n", err)
		return nil, err
	}
	return req, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
