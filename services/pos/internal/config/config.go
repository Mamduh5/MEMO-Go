package config

import (
	"os"
	"strconv"
	"time"
)

type MySQLConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

type JWTConfig struct {
	Secret          string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

type ServerConfig struct {
	POSGRPCPort  string
	AUTHGRPCPort string
}

type Config struct {
	MySQL  MySQLConfig
	JWT    JWTConfig
	Server ServerConfig
}

func Load() *Config {
	return &Config{
		MySQL: MySQLConfig{
			User:     getEnv("MYSQL_USER", "root"),
			Password: getEnv("MYSQL_PASSWORD", "password"),
			Host:     getEnv("MYSQL_HOST", "127.0.0.1"),
			Port:     getEnv("MYSQL_PORT", "3306"),
			DBName:   getEnv("MYSQL_DB", "memo_pos"),
		},
		JWT: JWTConfig{
			Secret:          getEnv("JWT_SECRET", "dev-secret-change-later"),
			AccessTokenTTL:  getDurationEnv("JWT_ACCESS_TTL_MIN", 15*time.Minute),
			RefreshTokenTTL: getDurationEnv("JWT_REFRESH_TTL_HOUR", 7*24*time.Hour),
		},
		Server: ServerConfig{
			POSGRPCPort:  getEnv("POS_GRPC_PORT", "50052"),
			AUTHGRPCPort: getEnv("AUTH_GRPC_PORT", "50051"),
		},
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getDurationEnv(key string, defaultVal time.Duration) time.Duration {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return time.Duration(i) * time.Minute
		}
	}
	return defaultVal
}
