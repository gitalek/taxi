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

	app := server.NewApp()
	// todo приведение типа?
	port := viper.GetString("port")
	if err := app.Run(port); err != nil {
		log.Fatalf("error while running server: %#v", err.Error())
	}
}
