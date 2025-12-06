package router

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/go-chi/chi/v5"

	"api-gateway/internal/config"
	"api-gateway/internal/handler"
	"api-gateway/internal/middleware"
)

func New(cfg config.Config) http.Handler {
	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.Logging)
	r.Use(middleware.CORS(cfg.AllowedOrigins))

	public := map[string]struct{}{
		"/api/v1/auth/register": {},
		"/api/v1/auth/login":    {},
		"/health":               {},
	}
	r.Use(middleware.Auth(cfg.JWTSecret, public))

	// routes
	r.Get("/health", handler.Health)

	authProxy := mustReverseProxy(cfg.AuthServiceURL)
	medicalProxy := mustReverseProxy(cfg.MedicalServiceURL)

	r.Mount("/api/v1/auth", authProxy)
	r.Mount("/api/v1/medical", medicalProxy)

	return r
}

// mustReverseProxy creates a reverse proxy to target; panics on error to fail fast.
func mustReverseProxy(target string) http.Handler {
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("invalid proxy target %s: %v", target, err)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		// preserve original host header for backend logging if needed
		req.Host = targetURL.Host
		// keep Authorization header as-is
	}
	proxy.ModifyResponse = func(resp *http.Response) error {
		// Drop downstream CORS headers to avoid duplicates with gateway CORS middleware.
		resp.Header.Del("Access-Control-Allow-Origin")
		resp.Header.Del("Access-Control-Allow-Credentials")
		resp.Header.Del("Access-Control-Allow-Headers")
		resp.Header.Del("Access-Control-Allow-Methods")
		return nil
	}
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("proxy error: %v", err)
		status := http.StatusBadGateway
		if strings.Contains(err.Error(), "timeout") {
			status = http.StatusGatewayTimeout
		}
		http.Error(w, "service unavailable", status)
	}

	return proxy
}
