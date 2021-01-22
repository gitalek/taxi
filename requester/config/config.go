package config

import "github.com/spf13/viper"

const configPath = "./requester/config"
const configName = "config"
const configType = "yaml"

func Init() error {
	viper.SetEnvPrefix("taxi")
	// TAXI_ORSKEY os var
	err := viper.BindEnv("orskey")
	if err != nil {
		return err
	}
	// TAXI_BINGMPKEY os var
	err = viper.BindEnv("bingmpkey")
	if err != nil {
		return err
	}

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	return viper.ReadInConfig()
}
