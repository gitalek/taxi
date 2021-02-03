package _map

import (
	"context"
	"github.com/gitalek/taxi/requester/pkg/types"
	"github.com/spf13/viper"
)

// todo should return error
func InitMaps() map[string]types.Requester {
	maps := make(map[string]types.Requester)
	// todo: dynamic?
	ors_url := viper.GetString("apiUrl")
	ors_key := viper.GetString("orskey")
	maps["ors"] = genRequestMetrics(ORSMetrics, ors_key, ors_url)
	bing_map_url := viper.GetString("bingMapsApi")
	bing_map_key := viper.GetString("bingmpkey")
	maps["bing_maps"] = genRequestMetrics(BingMapsMetrics, bing_map_key, bing_map_url)
	return maps
}

func genRequestMetrics(f types.RequestMetrics, key string, url string) types.Requester {
	return func(ctx context.Context, points []types.Point) (float64, float64, error) {
		return f(ctx, points, key, url)
	}
}
