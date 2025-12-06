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
	ErrRecordNotFound = errors.New("record not found")
)

type MedicalService struct {
	repo *repository.MedicalRepository
}

func NewMedicalService(repo *repository.MedicalRepository) *MedicalService {
	return &MedicalService{repo: repo}
}

func (s *MedicalService) Create(ctx context.Context, userID string, rec *models.MedicalRecord) error {
	rec.UserID = userID
	if rec.Title == "" {
		return fmt.Errorf("title required")
	}
	if rec.RecordDate.IsZero() {
		rec.RecordDate = time.Now()
	}
	return s.repo.Create(ctx, rec)
}

func (s *MedicalService) Get(ctx context.Context, userID, id string) (*models.MedicalRecord, error) {
	rec, err := s.repo.GetByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}
	if rec == nil {
		return nil, ErrRecordNotFound
	}
	return rec, nil
}

func (s *MedicalService) List(ctx context.Context, userID string, limit, offset int) ([]models.MedicalRecord, error) {
	if limit <= 0 {
		limit = 20
	}
	return s.repo.List(ctx, userID, limit, offset)
}

func (s *MedicalService) Update(ctx context.Context, userID string, rec *models.MedicalRecord) error {
	rec.UserID = userID
	if rec.Title == "" {
		return fmt.Errorf("title required")
	}
	return s.repo.Update(ctx, rec)
}

func (s *MedicalService) Delete(ctx context.Context, userID, id string) error {
	err := s.repo.Delete(ctx, id, userID)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrRecordNotFound
	}
	return err
}
