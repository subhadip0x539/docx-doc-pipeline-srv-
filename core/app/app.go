package app

import (
	"docx-doc-pipeline-srv/config"
	v1 "docx-doc-pipeline-srv/core/controller/rest/v1/router"

	"fmt"

	"github.com/gin-gonic/gin"
)

func Run(cfg config.Config) {
	handler := gin.Default()
	v1.NewRouter(handler)

	handler.Run(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port))
}
