package main

import (
	"fmt"
	calc "github.com/gitalek/taxi/calc/pkg"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	l "log"
	"net/http"
)

const port = "9090"

func main() {
	var svc calc.Service
	svc = &calc.CalcService{}
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	logger = log.WithPrefix(logger, "app", "calc", "layer", "logic")
	svc = calc.AppLoggingMiddleware{
		Logger: logger,
		Next:   svc,
	}

	logger = log.NewLogfmtLogger(log.StdlibWriter{})
	calculatePrice := calc.MakeCalculatePriceEndpoint(svc)
	calculatePrice = calc.LoggingMiddleware(
		log.WithPrefix(logger, "app", "calc", "layer", "transport: endpoint", "method", "calculatePrice"),
	)(calculatePrice)

	r := mux.NewRouter()
	r.Methods("POST").
		Path("/calcprice").
		Handler(
			httptransport.NewServer(
				calculatePrice,
				calc.DecodeRequest,
				calc.EncodeResponse,
			),
		)

	r.Methods("POST").
		Path("/v2/calcprice").
		Handler(
			httptransport.NewServer(
				calculatePrice,
				calc.DecodeRequestV2,
				calc.EncodeResponse,
			),
		)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "Taxi application")
		if err != nil {
			fmt.Printf("calc: handler: '/': error while writing response: %#v\n", err)
		}
	})

	address := fmt.Sprintf(":%s", port)
	l.Printf("Starting server at port %s\n", port)
	http.Handle("/", r)
	l.Fatal(http.ListenAndServe(address, nil))
}
