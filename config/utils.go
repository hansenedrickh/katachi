package config

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
)

func checkKey(key string) {
	if !viper.IsSet(key) && os.Getenv(key) == "" {
		panic(key + " key is not set")
	}
}

func FatalGetString(key string) string {
	checkKey(key)
	value := os.Getenv(key)
	if value == "" {
		value = viper.GetString(key)
	}
	return value
}

func GetIntOrPanic(key string) int {
	checkKey(key)
	v, err := strconv.Atoi(FatalGetString(key))
	if err != nil {
		v, err = strconv.Atoi(os.Getenv(key))
		panic("Could not parse key " + key)
	}
	return v
}
