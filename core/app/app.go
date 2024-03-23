package app

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"log/slog"

	"docx-doc-pipeline-srv/config"
	v1 "docx-doc-pipeline-srv/core/controller/rest/v1/router"
	"docx-doc-pipeline-srv/core/repo"
	"docx-doc-pipeline-srv/core/usecase"
	"docx-doc-pipeline-srv/pkg/mongo"
	"docx-doc-pipeline-srv/pkg/rabbit"
)

func Run(cfg config.Config) {
	db := mongo.NewMongo(cfg.Mongo.URI, int64(cfg.Mongo.Timeout))
	if err := db.Connect(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	client := db.GetClient()
	defer db.Disconnect()

	amqp := rabbit.NewRabbit(cfg.Rabbit.URI, int64(cfg.Rabbit.Timeout))
	if err := amqp.Connect(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	channel := amqp.GetChannel()
	defer amqp.Disconnect()

	uc := usecase.NewPipelineUseCase(repo.NewMongoRepo(client, cfg.Mongo.Database), repo.NewRabbitRepo(channel, cfg.Rabbit.Exchange))

	handler := gin.Default()
	v1.NewRouter(handler, uc)

	handler.Run(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port))
}
