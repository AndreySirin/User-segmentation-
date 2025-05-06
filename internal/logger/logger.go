package logger

import (
	"log/slog"
	"os"
)

type MyloggerInterface interface {
	Error(msg string, arg ...any)
	Info(msg string, arg ...any)
}

type MyLogger struct {
	lg *slog.Logger
}

func NewLogger() *MyLogger {
	return &MyLogger{
		lg: slog.New(slog.NewJSONHandler(os.Stdout, nil))}
}

func (m *MyLogger) Error(msg string, arg ...any) {
	m.lg.Error(msg, arg...)
}
func (m *MyLogger) Info(msg string, arg ...any) {
	m.lg.Info(msg, arg...)
}
