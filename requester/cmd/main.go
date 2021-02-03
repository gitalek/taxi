package main

import (
	"github.com/gitalek/taxi/requester/config"
	_map "github.com/gitalek/taxi/requester/pkg/map"
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

	config := server.AppConfig{
		Port: viper.GetString("port"),
		Maps: _map.InitMaps(),
	}
	app := server.NewApp(config)
	if err := app.Run(); err != nil {
		log.Fatalf("error while running server: %#v", err.Error())
	}
}
