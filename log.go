package mediatr

import (
	"log/slog"
	"os"
)

var (
	logger   *slog.Logger
	logLevel *slog.LevelVar
)

func init() {
	logLevel = new(slog.LevelVar)
	logLevel.Set(slog.LevelInfo)

	logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		AddSource:   false,
		Level:       logLevel,
		ReplaceAttr: nil,
	}))
}

func SetLogger(l *slog.Logger) {
	logger = l
}

func SetLogLevel(l slog.Level) {
	logLevel.Set(l)
}
