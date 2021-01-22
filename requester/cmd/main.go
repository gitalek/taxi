package main

import (
	"github.com/gitalek/taxi/requester/config"
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
		//todo приведение типа?
		//todo проверить на пустые поля
		Port:   viper.GetString("port"),
		ApiUrl: viper.GetString("apiUrl"),
		ORSKey: viper.GetString("orskey"),
	}
	app := server.NewApp(config)
	if err := app.Run(); err != nil {
		log.Fatalf("error while running server: %#v", err.Error())
	}
}
