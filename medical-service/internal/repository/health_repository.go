package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"medical-service/internal/models"
)

type HealthRepository struct {
	db DB
}

func NewHealthRepository(db DB) *HealthRepository {
	return &HealthRepository{db: db}
}

func (r *HealthRepository) Create(ctx context.Context, m *models.HealthMetric) error {
	const q = `
INSERT INTO health_metrics (user_id, metric_type, value, unit, recorded_at, notes)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, created_at;
`
	return r.db.QueryRowContext(ctx, q, m.UserID, m.MetricType, m.Value, m.Unit, m.RecordedAt, m.Notes).
		Scan(&m.ID, &m.CreatedAt)
}

func (r *HealthRepository) GetByID(ctx context.Context, id, userID string) (*models.HealthMetric, error) {
	const q = `
SELECT id, user_id, metric_type, value, unit, recorded_at, notes, created_at
FROM health_metrics
WHERE id = $1 AND user_id = $2;
`
	var m models.HealthMetric
	err := r.db.QueryRowContext(ctx, q, id, userID).Scan(
		&m.ID, &m.UserID, &m.MetricType, &m.Value, &m.Unit, &m.RecordedAt, &m.Notes, &m.CreatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("get metric: %w", err)
	}
	return &m, nil
}

func (r *HealthRepository) List(ctx context.Context, userID string, limit, offset int) ([]models.HealthMetric, error) {
	const q = `
SELECT id, user_id, metric_type, value, unit, recorded_at, notes, created_at
FROM health_metrics
WHERE user_id = $1
ORDER BY recorded_at DESC
LIMIT $2 OFFSET $3;
`
	rows, err := r.db.QueryContext(ctx, q, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list metrics: %w", err)
	}
	defer rows.Close()

	var res []models.HealthMetric
	for rows.Next() {
		var m models.HealthMetric
		if err := rows.Scan(&m.ID, &m.UserID, &m.MetricType, &m.Value, &m.Unit, &m.RecordedAt, &m.Notes, &m.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan metric: %w", err)
		}
		res = append(res, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *HealthRepository) Update(ctx context.Context, m *models.HealthMetric) error {
	const q = `
UPDATE health_metrics
SET metric_type=$1, value=$2, unit=$3, recorded_at=$4, notes=$5
WHERE id = $6 AND user_id = $7
RETURNING recorded_at;
`
	return r.db.QueryRowContext(ctx, q, m.MetricType, m.Value, m.Unit, m.RecordedAt, m.Notes, m.ID, m.UserID).
		Scan(&m.RecordedAt)
}

func (r *HealthRepository) Delete(ctx context.Context, id, userID string) error {
	const q = `DELETE FROM health_metrics WHERE id = $1 AND user_id = $2;`
	res, err := r.db.ExecContext(ctx, q, id, userID)
	if err != nil {
		return fmt.Errorf("delete metric: %w", err)
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
