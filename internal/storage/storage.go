package storage

import (
	"database/sql"
	"fmt"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/logger"
	_ "github.com/jackc/pgx/v5/stdlib"
	"net/url"
)

type Storage struct {
	lg *logger.MyLogger
	db *sql.DB
}

func New(log *logger.MyLogger, user, password, dbname, address string) (*Storage, error) {
	dsn := (&url.URL{
		Scheme: "postgresql",
		User:   url.UserPassword(user, password),
		Host:   address,
		Path:   dbname,
	}).String()

	log.Info(fmt.Sprintf("Подключение к базе: %s", dsn))

	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("connect to database: %w", err)
	}
	if err = sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}
	return &Storage{
		lg: log,
		db: sqlDB,
	}, nil
}

func (s *Storage) Close() error { return s.db.Close() }
