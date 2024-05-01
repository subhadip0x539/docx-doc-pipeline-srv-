package app

import (
	"log/slog"

	"docx-doc-pipeline-srv/src/config"
	"docx-doc-pipeline-srv/src/internal/adapter"
	"docx-doc-pipeline-srv/src/internal/core/repo"
	"docx-doc-pipeline-srv/src/internal/core/service"
	"docx-doc-pipeline-srv/src/internal/infra/http"
	"docx-doc-pipeline-srv/src/pkg/rabbit"
)

func Run(cfg config.Config) {
	amqp := rabbit.NewRabbit(cfg.AMQP.URI, int64(cfg.AMQP.Timeout))
	if err := amqp.Connect(); err != nil {
		slog.Error(err.Error())
	}
	amqpChannel := amqp.GetChannel()
	if err := amqpChannel.ExchangeDeclare(cfg.AMQP.Exchange, "topic", true, false, false, false, nil); err != nil {
		slog.Error(err.Error())
	}

	pipelineRepo := repo.NewPipelineRepo(amqpChannel, cfg.AMQP.Exchange)
	pipelineService := service.NewPipelineService(pipelineRepo)
	pipelineHandler := adapter.NewPipelineHandler(pipelineService)

	r, err := http.NewRouter(*pipelineHandler)
	if err != nil {
		slog.Error(err.Error())
	}
	s, err := http.NewServer(cfg.Server.Host, cfg.Server.Port, r)
	if err != nil {
		slog.Error(err.Error())
	}

	s.Serve()
}
