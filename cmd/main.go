package main

import (
	"context"
	"errors"
	"fmt"
	"os/signal"

	"golang.org/x/sync/errgroup"

	"github.com/AndreySirin/avito-backend-assignment-2023/internal/logger"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/server"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/service"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/storage"
)

// FIXME - убрать в конфиг файл
const (
	dbname   = "postgres"
	user     = "postgres"
	password = "secret"
	address  = "localhost:5432"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	ctx, cancel := signal.NotifyContext(context.Background())
	defer cancel()

	// TODO флаг debug брать из конфиг файла
	lg := logger.New(true)

	db, err := storage.New(lg, user, password, dbname, address)
	if err != nil {
		return fmt.Errorf("new database: %v", err)
	}
	defer func() {
		if errClose := db.Close(); errClose != nil {
			lg.With("error", errClose).Error("db.Close() in run()")
		}
	}()

	s := service.New(lg, db)

	srv := server.New(lg, ":8081", s)

	eg, ctx := errgroup.WithContext(ctx)

	// Run servers.
	eg.Go(func() error { return srv.Run(ctx) })

	if err = eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return fmt.Errorf("run: %v", err)
	}

	return nil
}
