package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"medical-service/internal/config"
	"medical-service/internal/handler"
	"medical-service/internal/middleware"
	"medical-service/internal/repository"
	"medical-service/internal/service"
)

func New(cfg config.Config, db repository.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logging)
	r.Use(middleware.CORS([]string{"*"}))

	medRepo := repository.NewMedicalRepository(db)
	healthRepo := repository.NewHealthRepository(db)

	medSvc := service.NewMedicalService(medRepo)
	healthSvc := service.NewHealthService(healthRepo)

	medH := handler.NewMedicalHandler(medSvc)
	healthH := handler.NewHealthHandler(healthSvc)

	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth(cfg.JWTSecret))

		r.Route("/api/v1/medical/records", func(r chi.Router) {
			r.Post("/", medH.Create)
			r.Get("/", medH.List)
			r.Get("/{id}", func(w http.ResponseWriter, req *http.Request) {
				medH.Get(w, req, chi.URLParam(req, "id"))
			})
			r.Put("/{id}", func(w http.ResponseWriter, req *http.Request) {
				medH.Update(w, req, chi.URLParam(req, "id"))
			})
			r.Delete("/{id}", func(w http.ResponseWriter, req *http.Request) {
				medH.Delete(w, req, chi.URLParam(req, "id"))
			})
		})

		r.Route("/api/v1/medical/health-metrics", func(r chi.Router) {
			r.Post("/", healthH.Create)
			r.Get("/", healthH.List)
			r.Get("/{id}", func(w http.ResponseWriter, req *http.Request) {
				healthH.Get(w, req, chi.URLParam(req, "id"))
			})
			r.Put("/{id}", func(w http.ResponseWriter, req *http.Request) {
				healthH.Update(w, req, chi.URLParam(req, "id"))
			})
			r.Delete("/{id}", func(w http.ResponseWriter, req *http.Request) {
				healthH.Delete(w, req, chi.URLParam(req, "id"))
			})
		})
	})

	r.Get("/health", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusOK) })

	return r
}
