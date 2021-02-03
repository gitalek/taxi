// setting up and running server
package server

import (
	"fmt"
	calc "github.com/gitalek/taxi/calc/pkg"
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
type calcApp struct {
	config AppConfig
}

// check interface realization
var _ App = &calcApp{}

type AppConfig struct {
	Port             string
	ApiUrl           string
	TaxiServicePrice float64
	MinPrice         float64
	MinuteRate       float64
	MeterRate        float64
}

// todo почему нельзя *App ?
func NewApp(config AppConfig) *calcApp {
	return &calcApp{config: config}
}

func (a calcApp) Run() error {
	var svc calc.Service
	serviceConfig := calc.ServiceConfig{
		ApiUrl:           a.config.ApiUrl,
		TaxiServicePrice: a.config.TaxiServicePrice,
		MinPrice:         a.config.MinPrice,
		MinuteRate:       a.config.MinuteRate,
		MeterRate:        a.config.MeterRate,
	}
	svc = calc.NewCalcService(serviceConfig)
	sugar := zap.NewExample().Sugar().With("app", "calc")
	defer func() {
		err := sugar.Sync()
		if err != nil {
			log.Fatalf("error while fleshing zap buffer: %#v\n", err)
		}
	}()
	svc = calc.AppLoggingMiddleware{
		Logger: sugar.With("layer", "logic"),
		Next:   svc,
	}

	calculatePrice := calc.MakeCalculatePriceEndpoint(svc)
	calculatePrice = calc.LoggingMiddleware(
		sugar.With("app", "calc", "layer", "transport: endpoint", "method", "calculatePrice"),
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

	address := fmt.Sprintf(":%s", a.config.Port)
	log.Printf("Starting server at port %s\n", a.config.Port)
	http.Handle("/", r)

	return http.ListenAndServe(address, nil)
}
