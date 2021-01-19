package config

import "github.com/spf13/viper"

const configPath = "./requester/config"
const configName = "config"
const configType = "yaml"

func Init() error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	return viper.ReadInConfig()
}
