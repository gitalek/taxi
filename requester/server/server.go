package server

import (
	"fmt"
	requester "github.com/gitalek/taxi/requester/pkg"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	l "log"
	"net/http"
)

type App interface {
	Run(string) error
}

// implementation of the interface
type requesterApp struct{}
// check interface realization
var _ App = &requesterApp{}

// todo почему нельзя *App ?
func NewApp() *requesterApp {
	return &requesterApp{}
}

func (a requesterApp) Run(port string) error {
	var svc requester.Service
	svc = &requester.RequesterService{}
	sugar := zap.NewExample().Sugar().With("app", "requester")
	defer sugar.Sync()
	svc = requester.AppLoggingMiddleware{
		Logger: sugar.With("layer", "logic"),
		Next:   svc,
	}

	tripMetrics := requester.MakeTripMetricsEndpoint(svc)
	tripMetrics = requester.LoggingMiddleware(
		sugar.With("app", "requester", "layer", "transport: endpoint", "method", "TripMetrics"),
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
	return server.ListenAndServe()
}
