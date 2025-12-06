package repository

import "context"

// DB is the minimal interface of *sql.DB used by repositories.
type DB interface {
	QueryRowContext(ctx context.Context, query string, args ...any) RowScanner
}

type RowScanner interface {
	Scan(dest ...any) error
}
