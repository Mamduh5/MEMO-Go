package config

import (
	"os"
)

type Config struct {
	POSGRPCPort  string
	JWTSecret    string
	AuthGRPCAddr string
}

func Load() (*Config, error) {
	return &Config{
		POSGRPCPort:  getEnv("POS_GRPC_PORT", "50052"),
		JWTSecret:    getEnv("JWT_SECRET", "dev-secret-change-later"),
		AuthGRPCAddr: getEnv("AUTH_GRPC_ADDR", "127.0.0.1:50051"),
	}, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
