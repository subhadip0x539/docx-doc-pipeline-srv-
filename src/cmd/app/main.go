package main

import (
	"os"

	"log/slog"

	"docx-doc-pipeline-srv/src/config"
	"docx-doc-pipeline-srv/src/internal/app"
	"docx-doc-pipeline-srv/src/pkg/logger"
	"docx-doc-pipeline-srv/src/pkg/motd"
)

func init() {
	motd.Info()
	logger.InitializeLogger()
}

func main() {
	if err := config.Register(".env", "env", os.Getenv("GIN_MODE")); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	cfg := config.GetConfig()

	app.Run(cfg)
}
