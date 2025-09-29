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

	// // *CREATE*
	historicoGroup.POST("/", handlers.CreateHistoricoHandler)
	// *GETALL*
	historicoGroup.GET("/", handlers.GetAllHistoricoHandler)
	// *GETBYID*
	historicoGroup.GET("/:id", handlers.GetHistoricoByIdHandler)
}
