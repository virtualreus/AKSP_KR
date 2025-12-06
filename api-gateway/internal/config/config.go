package config

import (
	"os"
)

type Config struct {
	Port            string
	AuthServiceURL  string
	MedicalServiceURL string
	JWTSecret       string
	AllowedOrigins  []string
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func Load() Config {
	return Config{
		Port:             getenv("GATEWAY_PORT", "8080"),
		AuthServiceURL:   getenv("AUTH_SERVICE_URL", "http://auth-service:8081"),
		MedicalServiceURL: getenv("MEDICAL_SERVICE_URL", "http://medical-service:8082"),
		JWTSecret:        getenv("JWT_SECRET", "dev-secret"),
		AllowedOrigins:   []string{getenv("CORS_ORIGIN", "*")},
	}
}

