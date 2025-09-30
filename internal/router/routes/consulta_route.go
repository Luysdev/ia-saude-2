package routes

import (
	"saude/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RouteConsulta(router *gin.Engine) {
	consultaGroup := router.Group("/consulta")

	// Endpoint de ping
	consultaGroup.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	consultaGroup.POST("/", handlers.CreateConsultaHandler)
	consultaGroup.GET("/", handlers.GetAllConsultaHandler)
	consultaGroup.GET("/:id", handlers.GetConsultaByIdHandler)
	consultaGroup.PUT("/:id", handlers.UpdateConsultaHandler)
	consultaGroup.DELETE("/:id", handlers.DeleteConsultaHandler)
}
