package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerPort  string
	DatabaseURL string
	JWTSecret   string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{
		ServerPort:  viper.GetString("SERVER_PORT"),
		DatabaseURL: viper.GetString("DATABASE_URL"),
		JWTSecret:   viper.GetString("JWT_SECRET"),
	}, nil
}
