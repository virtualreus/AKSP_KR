package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"medical-service/internal/middleware"
	"medical-service/internal/models"
	"medical-service/internal/service"
)

type HealthHandler struct {
	svc *service.HealthService
}

func NewHealthHandler(svc *service.HealthService) *HealthHandler {
	return &HealthHandler{svc: svc}
}

type metricRequest struct {
	MetricType string    `json:"metric_type"`
	Value      float64   `json:"value"`
	Unit       string    `json:"unit"`
	RecordedAt time.Time `json:"recorded_at"`
	Notes      string    `json:"notes"`
}

func (h *HealthHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID, _ := middleware.UserIDFromContext(r.Context())
	if userID == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	var req metricRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	m := &models.HealthMetric{
		MetricType: req.MetricType,
		Value:      req.Value,
		Unit:       req.Unit,
		RecordedAt: req.RecordedAt,
		Notes:      req.Notes,
	}
	if err := h.svc.Create(r.Context(), userID, m); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSON(w, http.StatusCreated, m)
}

func (h *HealthHandler) Get(w http.ResponseWriter, r *http.Request, id string) {
	userID, _ := middleware.UserIDFromContext(r.Context())
	if userID == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	m, err := h.svc.Get(r.Context(), userID, id)
	if err != nil {
		status := http.StatusInternalServerError
		if err == service.ErrMetricNotFound {
			status = http.StatusNotFound
		}
		http.Error(w, err.Error(), status)
		return
	}
	writeJSON(w, http.StatusOK, m)
}

func (h *HealthHandler) List(w http.ResponseWriter, r *http.Request) {
	userID, _ := middleware.UserIDFromContext(r.Context())
	if userID == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	limit := parseIntDefault(r.URL.Query().Get("limit"), 20)
	offset := parseIntDefault(r.URL.Query().Get("offset"), 0)

	metrics, err := h.svc.List(r.Context(), userID, limit, offset)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, metrics)
}

func (h *HealthHandler) Update(w http.ResponseWriter, r *http.Request, id string) {
	userID, _ := middleware.UserIDFromContext(r.Context())
	if userID == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	var req metricRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	m := &models.HealthMetric{
		ID:         id,
		MetricType: req.MetricType,
		Value:      req.Value,
		Unit:       req.Unit,
		RecordedAt: req.RecordedAt,
		Notes:      req.Notes,
	}
	if err := h.svc.Update(r.Context(), userID, m); err != nil {
		status := http.StatusBadRequest
		if err == service.ErrMetricNotFound {
			status = http.StatusNotFound
		}
		http.Error(w, err.Error(), status)
		return
	}
	writeJSON(w, http.StatusOK, m)
}

func (h *HealthHandler) Delete(w http.ResponseWriter, r *http.Request, id string) {
	userID, _ := middleware.UserIDFromContext(r.Context())
	if userID == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	if err := h.svc.Delete(r.Context(), userID, id); err != nil {
		status := http.StatusInternalServerError
		if err == service.ErrMetricNotFound {
			status = http.StatusNotFound
		}
		http.Error(w, err.Error(), status)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
