package _map

import (
	"context"
	"github.com/gitalek/taxi/requester/pkg/types"
	"net/http"
)

// todo should return error
func InitMaps(config types.MapsConfig) map[string]types.Requester {
	maps := make(map[string]types.Requester)
	//todo: dynamic?
	//todo: check property access errors
	orsConfig := config["ors"]
	bingConfig := config["bing"]
	maps["ors"] = genRequestMetrics(ORSMetrics, orsConfig.Token, orsConfig.Url)
	maps["bing_maps"] = genRequestMetrics(BingMapsMetrics, bingConfig.Token, bingConfig.Url)
	return maps
}

func genRequestMetrics(f types.RequestMetrics, key string, url string) types.Requester {
	return func(ctx context.Context, points []types.Point, client *http.Client) (float64, float64, error) {
		return f(ctx, points, key, url, client)
	}
}
