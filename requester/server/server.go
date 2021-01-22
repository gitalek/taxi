// setting up and running server
package server

import (
	"fmt"
	requester "github.com/gitalek/taxi/requester/pkg"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"log"
	"net/http"
)

type App interface {
	Run() error
}

// implementation of the interface
type requesterApp struct {
	config AppConfig
}

// check interface realization
var _ App = &requesterApp{}

type AppConfig struct {
	Port   string
	ApiUrl string
}

// todo почему нельзя *App ?
func NewApp(config AppConfig) *requesterApp {
	return &requesterApp{config: config}
}

func (a requesterApp) Run() error {
	var svc requester.Service
	serviceConfig := requester.ServiceConfig{
		ApiUrl: a.config.ApiUrl,
	}
	svc = &requester.RequesterService{Config: serviceConfig}
	sugar := zap.NewExample().Sugar().With("app", "requester")
	defer func() {
		err := sugar.Sync()
		if err != nil {
			log.Fatalf("error while fleshing zap buffer: %#v\n", err)
		}
	}()
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
		Addr:    fmt.Sprintf(":%s", a.config.Port),
		Handler: r,
	}

	log.Printf("Starting server at port %s\n", a.config.Port)
	return server.ListenAndServe()
}
