package http

import (
	"github.com/gin-gonic/gin"

	"docx-doc-pipeline-srv/src/internal/adapter"
)

func NewRouter(pipelineHandler adapter.TPipelineHandler) (*gin.Engine, error) {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		pipeline := v1.Group("/pipeline")
		{
			pipeline.POST("/dispatch", pipelineHandler.Dispatch)
		}
	}

	return router, nil
}
