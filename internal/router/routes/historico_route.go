package routes

import (
	"saude/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RouteHistorico(router *gin.Engine) {
	historicoGroup := router.Group("/historico")

	// Endpoint de ping
	historicoGroup.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	historicoGroup.POST("/", handlers.CreateHistoricoHandler)
	historicoGroup.GET("/", handlers.GetAllHistoricoHandler)
	historicoGroup.GET("/:id", handlers.GetHistoricoByIdHandler)
	historicoGroup.PUT("/:id", handlers.UpdateHistoricoHandler)
	historicoGroup.DELETE("/:id", handlers.DeleteHistoricoHandler)
}
