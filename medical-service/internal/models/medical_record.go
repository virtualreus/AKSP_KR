package models

import "time"

type MedicalRecord struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	RecordDate  time.Time `json:"record_date"`
	DoctorName  string    `json:"doctor_name,omitempty"`
	Diagnosis   string    `json:"diagnosis,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

