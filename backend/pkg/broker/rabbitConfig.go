package broker

import "os"

type RabbitConfig struct {
	URL string
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
