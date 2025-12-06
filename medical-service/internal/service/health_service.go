package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"medical-service/internal/models"
	"medical-service/internal/repository"
)

var (
	ErrMetricNotFound = errors.New("metric not found")
)

type HealthService struct {
	repo *repository.HealthRepository
}

func NewHealthService(repo *repository.HealthRepository) *HealthService {
	return &HealthService{repo: repo}
}

func (s *HealthService) Create(ctx context.Context, userID string, m *models.HealthMetric) error {
	m.UserID = userID
	if m.MetricType == "" || m.Value <= 0 || m.Unit == "" {
		return fmt.Errorf("metric_type, value (>0) and unit required")
	}
	if m.RecordedAt.IsZero() {
		m.RecordedAt = time.Now()
	}
	return s.repo.Create(ctx, m)
}

func (s *HealthService) Get(ctx context.Context, userID, id string) (*models.HealthMetric, error) {
	m, err := s.repo.GetByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}
	if m == nil {
		return nil, ErrMetricNotFound
	}
	return m, nil
}

func (s *HealthService) List(ctx context.Context, userID string, limit, offset int) ([]models.HealthMetric, error) {
	if limit <= 0 {
		limit = 20
	}
	return s.repo.List(ctx, userID, limit, offset)
}

func (s *HealthService) Update(ctx context.Context, userID string, m *models.HealthMetric) error {
	m.UserID = userID
	if m.MetricType == "" || m.Value <= 0 || m.Unit == "" {
		return fmt.Errorf("metric_type, value (>0) and unit required")
	}
	return s.repo.Update(ctx, m)
}

func (s *HealthService) Delete(ctx context.Context, userID, id string) error {
	err := s.repo.Delete(ctx, id, userID)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrMetricNotFound
	}
	return err
}
