package logger

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/goNiki/subservice/internal/infrastructure/config"
	errorapp "github.com/goNiki/subservice/internal/models/errorApp"
)

type Logger struct {
	log *slog.Logger
}

func InitLogger(cfg config.Logger) (*Logger, error) {
	var level slog.Level

	switch cfg.Level() {
	case "local":
		level = slog.LevelDebug
	case "dev":
		level = slog.LevelInfo
	case "prod":
		level = slog.LevelWarn
	default:
		level = slog.LevelError
	}

	opts := &slog.HandlerOptions{
		Level: level,
	}
	var handlers []slog.Handler

	if cfg.Console().Enabled() {
		w := os.Stdout
		if cfg.Console().Output() == "stderr" {
			w = os.Stderr
		}
		switch cfg.Console().Format() {
		case "text":
			handlers = append(handlers, slog.NewTextHandler(w, opts))
		case "json":
			handlers = append(handlers, slog.NewJSONHandler(w, opts))
		default:
			handlers = append(handlers, slog.NewJSONHandler(w, opts))
		}
	}

	if cfg.File().Enabled() {
		file, err := os.OpenFile(cfg.File().FilePath(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", errorapp.ErrFailedOpenLogFile, err)
		}
		switch cfg.File().Format() {
		case "text":
			handlers = append(handlers, slog.NewTextHandler(file, opts))
		case "json":
			handlers = append(handlers, slog.NewJSONHandler(file, opts))
		default:
			handlers = append(handlers, slog.NewJSONHandler(file, opts))
		}
	}

	handler := slog.NewMultiHandler(handlers...)

	logger := slog.New(handler)

	return &Logger{
		log: logger,
	}, nil
}
