package logger

import (
	"os"

	"log/slog"
)

func InitializeLogger() {
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})

	slog.SetDefault(slog.New(logHandler))
}
