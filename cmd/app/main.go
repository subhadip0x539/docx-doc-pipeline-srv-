package main

import (
	"docx-doc-pipeline-srv/config"
	"docx-doc-pipeline-srv/core/app"
	"docx-doc-pipeline-srv/pkg/logger"
	"docx-doc-pipeline-srv/pkg/motd"
	"log/slog"
	"os"
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
