package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/url"

	_ "github.com/jackc/pgx/v5/stdlib" // импорт драйвера pgx для database/sql
)

const module = "storage"

type Storage struct {
	lg *slog.Logger
	db *sql.DB
}

func New(lg *slog.Logger, user, password, dbname, address string) (*Storage, error) {
	dsn := (&url.URL{
		Scheme: "postgresql",
		User:   url.UserPassword(user, password),
		Host:   address,
		Path:   dbname,
	}).String()

	lg.Debug(dsn)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("connect to database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return &Storage{
		lg: lg.With("module", module),
		db: db,
	}, nil
}

func (s *Storage) Close() error { return s.db.Close() }

func (s *Storage) TX(ctx context.Context) (*sql.Tx, error) {
	return s.db.BeginTx(ctx, nil)
}
