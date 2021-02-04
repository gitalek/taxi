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

	maps := make(types.MapsConfig)
	// todo: dynamic?
	orsUrl := viper.GetString("apiUrl")
	orsToken := viper.GetString("orskey")
	bingUrl := viper.GetString("bingMapsApi")
	bingToken := viper.GetString("bingmpkey")

	maps["ors"] = types.MapConfig{Url: orsUrl, Token: orsToken}
	maps["bing"] = types.MapConfig{Url: bingUrl, Token: bingToken}

	config := server.AppConfig{
		Port: viper.GetString("port"),
		Maps: maps,
	}
	app := server.NewApp(config)
	if err := app.Run(); err != nil {
		log.Fatalf("error while running server: %#v", err.Error())
	}
}
