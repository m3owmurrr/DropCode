package config

import (
	"os"

	"github.com/m3owmurrr/dropcode/backend/pkg/broker"
	"github.com/m3owmurrr/dropcode/backend/pkg/storage"
)

type Config struct {
	Host   string
	Port   string
	S3     storage.S3Config
	Rabbit broker.RabbitConfig
}

var Cfg *Config

func Load() {
	Cfg = &Config{
		Host:   getEnv("HOST", ""),
		Port:   getEnv("PORT", "8080"),
		S3:     storage.LoadS3Config(),
		Rabbit: broker.LoadRabbitConfig(),
	}
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
