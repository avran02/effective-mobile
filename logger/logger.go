package logger

import (
	"log/slog"
	"os"

	"github.com/avran02/effective-mobile/config"
)

func Setup(conf config.Server) {
	var ll slog.Leveler
	switch conf.LogLevel {
	case "debug":
		ll = slog.LevelDebug
	case "info":
		ll = slog.LevelInfo
	case "warn":
		ll = slog.LevelWarn
	case "error":
		ll = slog.LevelError
	default:
		ll = slog.LevelInfo
	}

	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: ll,
	}))

	slog.SetDefault(log)
}
