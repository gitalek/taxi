package main

import (
	"github.com/gitalek/taxi/calc/config"
	"github.com/gitalek/taxi/calc/server"
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
		Port:        viper.GetString("port"),
		ApiUrl:      viper.GetString("apiUrl"),
		TaxiService: viper.GetInt("taxiService"),
		MinPrice:    viper.GetInt("minPrice"),
		MinuteRate:  viper.GetInt("minuteRate"),
		KmRate:      viper.GetInt("kmRate"),
	}

	app := server.NewApp(config)
	if err := app.Run(); err != nil {
		log.Fatalf("error while running server: %#v", err.Error())
	}
}
