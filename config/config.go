package config

import (
	"github.com/spf13/viper"

	"github.com/hansenedrickh/katachi/utils"
)

type Config struct {
	Port     int
	Keisatsu *KeisatsuConfig
}

var appConfig *Config

func Load() {
	viper.AutomaticEnv()
	viper.SetConfigName("application")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.SetConfigType("yaml")

	viper.ReadInConfig()

	appConfig = &Config{
		Port:     utils.GetIntOrPanic("APP_PORT"),
		Keisatsu: newKeisatsuConfig(),
	}
}

func Port() int {
	return appConfig.Port
}

func Keisatsu() *KeisatsuConfig {
	return appConfig.Keisatsu
}
