package router

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"

	"auth-service/internal/config"
	"auth-service/internal/handler"
	"auth-service/internal/middleware"
	"auth-service/internal/repository"
	"auth-service/internal/service"
)

func New(cfg config.Config, dbConn *sql.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logging)
	r.Use(middleware.CORS([]string{"*"}))

	userRepo := repository.NewUserRepository(dbConn)
	authSvc := service.NewAuthService(userRepo, cfg.JWTSecret)
	h := handler.NewAuthHandler(authSvc)

	r.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/register", h.Register)
		r.Post("/login", h.Login)
		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(cfg.JWTSecret))
			r.Get("/me", h.Me)
		})
	})

	r.Get("/health", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusOK) })

	return r
}
