package config

import "os"

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
		Port:        getenv("MEDICAL_PORT", "8082"),
		DatabaseURL: getenv("MEDICAL_DATABASE_URL", "postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable"),
		JWTSecret:   getenv("JWT_SECRET", "dev-secret"),
	}
}
