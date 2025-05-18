package main

import (
	"context"
	"errors"
	"fmt"
	"os/signal"

	"golang.org/x/sync/errgroup"

	"github.com/AndreySirin/avito-backend-assignment-2023/internal/config"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/logger"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/server"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/service"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/storage"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	ctx, cancel := signal.NotifyContext(context.Background())
	defer cancel()

	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("error loading config: %v", err)
	}
	lg := logger.New(cfg.Logger.Debug)
	db, err := storage.New(
		lg,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DbName,
		cfg.Postgres.Address,
	)
	if err != nil {
		return fmt.Errorf("new database: %v", err)
	}
	defer func() {
		if errClose := db.Close(); errClose != nil {
			lg.With("error", errClose).Error("db.Close() in run()")
		}
	}()

	s := service.New(lg, db)

	srv := server.New(lg, cfg.Server.Port, s)

	eg, ctx := errgroup.WithContext(ctx)

	// Run servers.
	eg.Go(func() error { return srv.Run(ctx, cfg.Server.ShutdownTimeout) })

	if err = eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return fmt.Errorf("run: %v", err)
	}

	return nil
}
