package logger

import (
	"log/slog"
	"os"
	"time"

	"github.com/Nadzyoki/wowmq/internal/config"
	"github.com/lmittmann/tint"
)

const (
	version = "VERSION_LOGGER"
	local   = "LOCAL"
	prod    = "PROD"
)

func NewLogger(cfg *config.Config) *slog.Logger {
	ver := os.Getenv(version)
	switch ver {
	case local:
		w := os.Stderr

		logger := slog.New(tint.NewHandler(w, nil))

		slog.SetDefault(slog.New(
			tint.NewHandler(w, &tint.Options{
				Level:      slog.LevelDebug,
				TimeFormat: time.Kitchen,
			}),
		))
		return logger

	case prod:
	}
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)

	logger := slog.New(handler)

	return logger
}
