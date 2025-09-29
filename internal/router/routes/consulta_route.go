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

	// // *CREATE*
	consultaGroup.POST("/", handlers.CreateConsultaHandler)
	// *GETALL*
	consultaGroup.GET("/", handlers.GetAllConsultaHandler)
	// *GETBYID*
	consultaGroup.GET("/:id", handlers.GetConsultaByIdHandler)
}
