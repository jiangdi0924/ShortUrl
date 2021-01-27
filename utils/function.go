package utils

import (
	"log"

	"github.com/spf13/viper"
)

func Env(key string, fallback string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		if fallback != "" {
			return fallback
		}
		log.Fatalf("Invalid type assertion")
	}

	return value
}
