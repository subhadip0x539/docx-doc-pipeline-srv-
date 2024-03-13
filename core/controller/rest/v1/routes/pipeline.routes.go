package routes

import (
	"github.com/gin-gonic/gin"
)

func NewPLRoutes(rg *gin.RouterGroup) {

	handler := rg.Group("/pipeline")
	{
		handler.POST("/process")
	}
}
