package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"medical-service/internal/middleware"
	"medical-service/internal/models"
	"medical-service/internal/service"
)

type MedicalHandler struct {
	svc *service.MedicalService
}

func NewMedicalHandler(svc *service.MedicalService) *MedicalHandler {
	return &MedicalHandler{svc: svc}
}

type recordRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	RecordDate  time.Time `json:"record_date"`
	DoctorName  string    `json:"doctor_name"`
	Diagnosis   string    `json:"diagnosis"`
}

func (h *MedicalHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID, _ := middleware.UserIDFromContext(r.Context())
	if userID == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var req recordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	rec := &models.MedicalRecord{
		Title:       req.Title,
		Description: req.Description,
		RecordDate:  req.RecordDate,
		DoctorName:  req.DoctorName,
		Diagnosis:   req.Diagnosis,
	}
	if err := h.svc.Create(r.Context(), userID, rec); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusCreated, rec)
}

func (h *MedicalHandler) Get(w http.ResponseWriter, r *http.Request, id string) {
	userID, _ := middleware.UserIDFromContext(r.Context())
	if userID == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	rec, err := h.svc.Get(r.Context(), userID, id)
	if err != nil {
		status := http.StatusInternalServerError
		if err == service.ErrRecordNotFound {
			status = http.StatusNotFound
		}
		http.Error(w, err.Error(), status)
		return
	}
	writeJSON(w, http.StatusOK, rec)
}

func (h *MedicalHandler) List(w http.ResponseWriter, r *http.Request) {
	userID, _ := middleware.UserIDFromContext(r.Context())
	if userID == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	limit := parseIntDefault(r.URL.Query().Get("limit"), 20)
	offset := parseIntDefault(r.URL.Query().Get("offset"), 0)

	recs, err := h.svc.List(r.Context(), userID, limit, offset)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, recs)
}

func (h *MedicalHandler) Update(w http.ResponseWriter, r *http.Request, id string) {
	userID, _ := middleware.UserIDFromContext(r.Context())
	if userID == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	var req recordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	rec := &models.MedicalRecord{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		RecordDate:  req.RecordDate,
		DoctorName:  req.DoctorName,
		Diagnosis:   req.Diagnosis,
	}
	if err := h.svc.Update(r.Context(), userID, rec); err != nil {
		status := http.StatusBadRequest
		if err == service.ErrRecordNotFound {
			status = http.StatusNotFound
		}
		http.Error(w, err.Error(), status)
		return
	}
	writeJSON(w, http.StatusOK, rec)
}

func (h *MedicalHandler) Delete(w http.ResponseWriter, r *http.Request, id string) {
	userID, _ := middleware.UserIDFromContext(r.Context())
	if userID == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	if err := h.svc.Delete(r.Context(), userID, id); err != nil {
		status := http.StatusInternalServerError
		if err == service.ErrRecordNotFound {
			status = http.StatusNotFound
		}
		http.Error(w, err.Error(), status)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
