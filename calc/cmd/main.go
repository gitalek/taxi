package main

import (
	"fmt"
	calc "github.com/gitalek/taxi/calc/pkg"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	l "log"
	"net/http"
	"os"
)

const port = "9090"

func main() {
	//ctx := context.Background()
	svc := &calc.CalcService{}
	logger := log.NewLogfmtLogger(os.Stderr)
	svc_with_log := &calc.AppLoggingMiddleware{logger, svc}

	calculatePrice := calc.MakeCalculatePriceEndpoint(svc_with_log)
	calculatePrice = calc.LoggingMiddleware(log.With(logger, "method", "calculatePrice"))(calculatePrice)
	//calculatePriceHandler := calc.MakeHttpHandler(ctx, calculatePrice)

	r := mux.NewRouter()
	r.Methods("POST").
		Path("/calcprice").
		Handler(
			httptransport.NewServer(
				calc.MakeCalculatePriceEndpoint(svc),
				calc.DecodeRequest,
				calc.EncodeResponse,
				//httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
			),
		)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello!")
	})

	r.HandleFunc("/trip", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "trip page!")
	})

	address := fmt.Sprintf(":%s", port)
	l.Printf("Starting server at port %s\n", port)
	http.Handle("/", r)
	http.ListenAndServe(address, nil)
}
