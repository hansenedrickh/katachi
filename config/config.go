package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port           int
	JWTSecretToken string
	Database       DatabaseConfig
}

type DatabaseConfig struct {
	Host     string `cfg:"DATABASE_HOST"`
	Port     string `cfg:"DATABASE_PORT"`
	Username string `cfg:"DATABASE_USERNAME"`
	Password string `cfg:"DATABASE_PASSWORD"`
	Name     string `cfg:"DATABASE_NAME"`
}

func Load() Config {
	viper.AutomaticEnv()
	viper.SetConfigName("application")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.SetConfigType("yaml")

	viper.ReadInConfig()

	return Config{
		Port:           GetIntOrPanic("APP_PORT"),
		JWTSecretToken: FatalGetString("JWT_SECRET_TOKEN"),
		Database: DatabaseConfig{
			Host:     FatalGetString("DATABASE_HOST"),
			Port:     FatalGetString("DATABASE_PORt"),
			Username: FatalGetString("DATABASE_USERNAME"),
			Password: FatalGetString("DATABASE_PASSWORD"),
			Name:     FatalGetString("DATABASE_NAME"),
		},
	}
}
