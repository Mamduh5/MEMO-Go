package config

import (
	"os"
)

type Config struct {
	GRPCPort  string
	JWTSecret string
}

func Load() (*Config, error) {
	return &Config{
		GRPCPort:  getEnv("POS_GRPC_PORT", "50052"),
		JWTSecret: getEnv("JWT_SECRET", ""),
	}, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
