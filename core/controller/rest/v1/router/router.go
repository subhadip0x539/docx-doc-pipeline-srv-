package router

import (
	"docx-doc-pipeline-srv/core/controller/rest/v1/routes"

	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine) {
	handler.GET("/health", func(c *gin.Context) { c.Status(http.StatusOK) })

	rg := handler.Group("/v1")
	{
		routes.NewPLRoutes(rg)
	}
}
