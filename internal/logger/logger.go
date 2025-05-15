package logger

import (
	"log/slog"
	"os"
)

func New(debug bool) *slog.Logger {
	lg := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelError,
	}))

	if debug {
		lg = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
		lg.Info("DEBUG MODE")
	}

	return lg
}
