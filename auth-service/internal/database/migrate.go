package database

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func Migrate(ctx context.Context, db *sql.DB) error {
	goose.SetBaseFS(embedMigrations)
	// Разделяем таблицу версий для общего Postgres, чтобы не конфликтовать с другими сервисами.
	goose.SetTableName("goose_db_version_auth")
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("set dialect: %w", err)
	}
	if err := goose.UpContext(ctx, db, "migrations"); err != nil {
		return fmt.Errorf("migrate: %w", err)
	}
	return nil
}
