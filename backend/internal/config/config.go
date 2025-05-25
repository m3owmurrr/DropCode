package config

import "os"

type Config struct {
	Port   string
	S3     S3Config
	Rabbit RabbitConfig
}

type S3Config struct {
	Endpoint   string
	AccessKey  string
	SecretKey  string
	Region     string
	RunBucket  string
	SaveBucket string
}

type RabbitConfig struct {
	URL string
}

var Cfg *Config

func Load() {
	Cfg = &Config{
		Port: getEnv("PORT", "8080"),
		S3: S3Config{
			Endpoint:   getEnv("S3_ENDPOINT", ""),
			AccessKey:  getEnv("S3_ACCESS_KEY", ""),
			SecretKey:  getEnv("S3_SECRET_KEY", ""),
			Region:     getEnv("S3_REGION", "us-east-1"),
			RunBucket:  getEnv("S3_RUN_BUCKET", ""),
			SaveBucket: getEnv("S3_SAVE_BUCKET", ""),
		},
		Rabbit: RabbitConfig{
			URL: getEnv("RABBIT_URL", ""),
		},
	}
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
