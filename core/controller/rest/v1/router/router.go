package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"docx-doc-pipeline-srv/core/controller/rest/v1/routes"
	"docx-doc-pipeline-srv/core/usecase"
)

func NewRouter(handler *gin.Engine, uc usecase.IPipelineUseCase) {
	handler.GET("/health", func(c *gin.Context) { c.Status(http.StatusOK) })

	rg := handler.Group("/v1")
	{
		routes.NewPipelineRoutes(rg, uc)
	}
}
