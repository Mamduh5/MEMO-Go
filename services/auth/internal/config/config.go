package config

import (
	"os"
)

type MySQLConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

type Config struct {
	MySQL MySQLConfig
}

func Load() *Config {
	return &Config{
		MySQL: MySQLConfig{
			User:     getEnv("MYSQL_USER", "root"),
			Password: getEnv("MYSQL_PASSWORD", "password"),
			Host:     getEnv("MYSQL_HOST", "mysql"),
			Port:     getEnv("MYSQL_PORT", "3306"),
			DBName:   getEnv("MYSQL_DB", "memo_auth"),
		},
	}
}

func getEnv(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}
