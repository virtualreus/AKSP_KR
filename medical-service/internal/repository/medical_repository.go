package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"medical-service/internal/models"
)

type MedicalRepository struct {
	db DB
}

func NewMedicalRepository(db DB) *MedicalRepository {
	return &MedicalRepository{db: db}
}

func (r *MedicalRepository) Create(ctx context.Context, rec *models.MedicalRecord) error {
	const q = `
INSERT INTO medical_records (user_id, title, description, record_date, doctor_name, diagnosis)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, created_at, updated_at;
`
	return r.db.QueryRowContext(ctx, q, rec.UserID, rec.Title, rec.Description, rec.RecordDate, rec.DoctorName, rec.Diagnosis).
		Scan(&rec.ID, &rec.CreatedAt, &rec.UpdatedAt)
}

func (r *MedicalRepository) GetByID(ctx context.Context, id, userID string) (*models.MedicalRecord, error) {
	const q = `
SELECT id, user_id, title, description, record_date, doctor_name, diagnosis, created_at, updated_at
FROM medical_records
WHERE id = $1 AND user_id = $2;
`
	var rec models.MedicalRecord
	err := r.db.QueryRowContext(ctx, q, id, userID).Scan(
		&rec.ID, &rec.UserID, &rec.Title, &rec.Description, &rec.RecordDate, &rec.DoctorName, &rec.Diagnosis, &rec.CreatedAt, &rec.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("get record: %w", err)
	}
	return &rec, nil
}

func (r *MedicalRepository) List(ctx context.Context, userID string, limit, offset int) ([]models.MedicalRecord, error) {
	const q = `
SELECT id, user_id, title, description, record_date, doctor_name, diagnosis, created_at, updated_at
FROM medical_records
WHERE user_id = $1
ORDER BY record_date DESC
LIMIT $2 OFFSET $3;
`
	rows, err := r.db.QueryContext(ctx, q, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list records: %w", err)
	}
	defer rows.Close()

	var res []models.MedicalRecord
	for rows.Next() {
		var rec models.MedicalRecord
		if err := rows.Scan(&rec.ID, &rec.UserID, &rec.Title, &rec.Description, &rec.RecordDate, &rec.DoctorName, &rec.Diagnosis, &rec.CreatedAt, &rec.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan record: %w", err)
		}
		res = append(res, rec)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *MedicalRepository) Update(ctx context.Context, rec *models.MedicalRecord) error {
	rec.UpdatedAt = time.Now()
	const q = `
UPDATE medical_records
SET title=$1, description=$2, record_date=$3, doctor_name=$4, diagnosis=$5, updated_at=$6
WHERE id = $7 AND user_id = $8
RETURNING updated_at;
`
	return r.db.QueryRowContext(ctx, q, rec.Title, rec.Description, rec.RecordDate, rec.DoctorName, rec.Diagnosis, rec.UpdatedAt, rec.ID, rec.UserID).
		Scan(&rec.UpdatedAt)
}

func (r *MedicalRepository) Delete(ctx context.Context, id, userID string) error {
	const q = `DELETE FROM medical_records WHERE id = $1 AND user_id = $2;`
	res, err := r.db.ExecContext(ctx, q, id, userID)
	if err != nil {
		return fmt.Errorf("delete record: %w", err)
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
