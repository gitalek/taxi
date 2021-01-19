package server

import (
	"fmt"
	calc "github.com/gitalek/taxi/calc/pkg"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	l "log"
	"net/http"
)

type App interface {
	Run(string) error
}

// implementation of the interface
type calcApp struct{}
// check interface realization
var _ App = &calcApp{}

// todo почему нельзя *App ?
func NewApp() *calcApp {
	return &calcApp{}
}

func (a calcApp) Run(port string) error {
	var svc calc.Service
	svc = &calc.CalcService{}
	loggerSvc := log.NewLogfmtLogger(log.StdlibWriter{})
	loggerSvc = log.WithPrefix(loggerSvc, "app", "calc", "layer", "logic")
	svc = calc.AppLoggingMiddleware{
		Logger: loggerSvc,
		Next:   svc,
	}

	loggerEndpoint := log.NewLogfmtLogger(log.StdlibWriter{})
	calculatePrice := calc.MakeCalculatePriceEndpoint(svc)
	calculatePrice = calc.LoggingMiddleware(
		log.WithPrefix(loggerEndpoint, "app", "calc", "layer", "transport: endpoint", "method", "calculatePrice"),
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

	return http.ListenAndServe(address, nil)
}
