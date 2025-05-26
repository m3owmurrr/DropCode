package config

import (
	"os"

	"github.com/m3owmurrr/dropcode/backend/pkg/config"
)

type Config struct {
	Host   string
	Port   string
	S3     config.S3Config
	Rabbit config.RabbitConfig
}

var Cfg *Config

func Load() {
	Cfg = &Config{
		Host:   getEnv("HOST", ""),
		Port:   getEnv("PORT", "8080"),
		S3:     config.LoadS3Config(),
		Rabbit: config.LoadRabbitConfig(),
	}
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
