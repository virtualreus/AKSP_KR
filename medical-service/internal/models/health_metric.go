package models

import "time"

type HealthMetric struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	MetricType string    `json:"metric_type"`
	Value      float64   `json:"value"`
	Unit       string    `json:"unit"`
	RecordedAt time.Time `json:"recorded_at"`
	Notes      string    `json:"notes,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}
