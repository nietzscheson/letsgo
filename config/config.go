package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	DatabaseURL   string `mapstructure:"DATABASE_URL"`
}

func LoadConfig() (config Config) {
	viper.AutomaticEnv()

	viper.BindEnv("SERVER_ADDRESS")
	viper.BindEnv("DATABASE_URL")

	viper.SetDefault("SERVER_ADDRESS", ":8080")

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return config
}
