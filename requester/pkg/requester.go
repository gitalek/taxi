package requester

import (
	"context"
	"github.com/gitalek/taxi/requester/pkg/strategies"
	"github.com/gitalek/taxi/requester/pkg/types"
)

func init() {

}

// service as an interface
type Service interface {
	TripMetrics(context.Context, []types.Point, int) (float64, float64, error)
}

type ServiceConfig struct {
	Maps map[string]types.Requester
}

// implementation of the interface
type RequesterService struct {
	Config ServiceConfig
}

// check interface realization
var _ Service = &RequesterService{}

func (s *RequesterService) TripMetrics(ctx context.Context, c []types.Point, strategy int) (float64, float64, error) {
	//todo: повторные вычисления!
	strtgs := strategies.InitStrategies()
	if strategy < 1 || strategy > len(strtgs) {
		strategy = 1
	}
	return strtgs[strategy-1](ctx, c, s.Config.Maps)
}
