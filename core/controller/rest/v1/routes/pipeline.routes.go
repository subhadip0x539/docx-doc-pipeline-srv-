package routes

import (
	"github.com/gin-gonic/gin"

	"docx-doc-pipeline-srv/core/usecase"
)

type PipelineRoutes struct {
	usecase usecase.IPipelineUseCase
}

func (r *PipelineRoutes) Documents(c *gin.Context) {
}

func NewPipelineRoutes(rg *gin.RouterGroup, uc usecase.IPipelineUseCase) {
	routes := &PipelineRoutes{usecase: uc}

	handler := rg.Group("/pipeline")
	{
		handler.POST("/documents", routes.Documents)
	}
}
