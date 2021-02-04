package config

import "github.com/spf13/viper"

const configPath = "./requester/config"
const configName = "config"
const configType = "yaml"

func Init() error {
	viper.SetEnvPrefix("taxi")
	// TAXI_ORS_TOKEN os var
	err := viper.BindEnv("ors_token")
	if err != nil {
		return err
	}
	// TAXI_BING_TOKEN os var
	err = viper.BindEnv("bing_token")
	if err != nil {
		return err
	}

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	return viper.ReadInConfig()
}
