package calc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func DecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Printf("%#v\n", err)
		return nil, err
	}
	return req, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

//func MakeHttpHandler(ctx context.Context, endpoint endpoint.Endpoint) http.Handler {
//	r := mux.NewRouter()
//	r.Methods("POST").
//		Path("/taxi/{x}/{y}").
//		Handler(
//			httptransport.NewServer(
//				endpoint,
//				decodeRequest,
//				encodeResponse,
//			),
//		)
//	return r
//}
