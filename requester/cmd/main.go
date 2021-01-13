package main

import (
	"fmt"
	requester "github.com/gitalek/taxi/requester/pkg"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const port = "9091"

func main() {
	svc := &requester.RequesterService{}

	r := mux.NewRouter()
	r.Methods("POST").
		Path("/tripmetrics").
		Handler(
			httptransport.NewServer(
				requester.MakeTripMetricsEndpoint(svc),
				requester.DecodeRequest,
				requester.EncodeResponse,
			),
		)

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	log.Fatal(server.ListenAndServe())
}
