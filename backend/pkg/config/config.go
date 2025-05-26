package config

import "os"

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

func LoadS3Config() S3Config {
	return S3Config{
		Endpoint:   getEnv("S3_ENDPOINT", ""),
		AccessKey:  getEnv("S3_ACCESS_KEY", ""),
		SecretKey:  getEnv("S3_SECRET_KEY", ""),
		Region:     getEnv("S3_REGION", ""),
		RunBucket:  getEnv("S3_RUN_BUCKET", ""),
		SaveBucket: getEnv("S3_SAVE_BUCKET", ""),
	}
}

func LoadRabbitConfig() RabbitConfig {
	return RabbitConfig{
		URL: getEnv("RABBIT_URL", ""),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
