package main

import (
	"fmt"
	requester "github.com/gitalek/taxi/requester/pkg"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	l "log"
	"net/http"
)

const port = "9091"

func main() {
	var svc requester.Service
	svc = &requester.RequesterService{}
	loggerSvc := log.NewLogfmtLogger(log.StdlibWriter{})
	loggerSvc = log.WithPrefix(loggerSvc, "app", "requester", "layer", "logic")
	svc = requester.AppLoggingMiddleware{
		Logger: loggerSvc,
		Next:   svc,
	}

	loggerEndpoint := log.NewLogfmtLogger(log.StdlibWriter{})
	tripMetrics := requester.MakeTripMetricsEndpoint(svc)
	tripMetrics = requester.LoggingMiddleware(
		log.WithPrefix(loggerEndpoint, "app", "requester", "layer", "transport: endpoint", "method", "TripMetrics"),
	)(tripMetrics)

	r := mux.NewRouter()
	r.Methods("POST").
		Path("/tripmetrics").
		Handler(
			httptransport.NewServer(
				tripMetrics,
				requester.DecodeRequest,
				requester.EncodeResponse,
			),
		)

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	l.Printf("Starting server at port %s\n", port)
	l.Fatal(server.ListenAndServe())
}
