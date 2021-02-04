package main

import (
	"github.com/gitalek/taxi/requester/config"
	"github.com/gitalek/taxi/requester/pkg/types"
	"github.com/gitalek/taxi/requester/server"
	"github.com/spf13/viper"
	"log"
)

func main() {
	// reading config data into viper
	err := config.Init()
	if err != nil {
		log.Fatalf("error while reading config: %#v\n", err)
	}

	appConfig := server.AppConfig{
		Port: viper.GetString("port"),
		Maps: InitMapsConfig(),
	}
	app := server.NewApp(appConfig)
	if err := app.Run(); err != nil {
		log.Fatalf("error while running server: %#v", err.Error())
	}
}

func InitMapsConfig() types.MapsConfig {
	maps := make(types.MapsConfig)
	//todo: dynamic?
	//todo: check access properties errors
	orsUrl := viper.GetString("orsUrl")
	orsToken := viper.GetString("ors_token")
	bingUrl := viper.GetString("bingUrl")
	bingToken := viper.GetString("bing_token")
	maps["ors"] = types.MapConfig{Url: orsUrl, Token: orsToken}
	maps["bing"] = types.MapConfig{Url: bingUrl, Token: bingToken}
	return maps
}
