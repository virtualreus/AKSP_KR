package config

import (
	"os"
)

type Config struct {
	Port        string
	DatabaseURL string
	JWTSecret   string
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func Load() Config {
	return Config{
		Port:        getenv("AUTH_PORT", "8081"),
		DatabaseURL: getenv("AUTH_DATABASE_URL", "postgres://postgres:postgres@postgres-auth:5432/postgres?sslmode=disable"),
		JWTSecret:   getenv("JWT_SECRET", "dev-secret"),
	}
}
