package routes

import (
	"saude/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RouteAnalyse(router *gin.Engine) {
	analyseGroup := router.Group("/analyse")

	// Endpoint de ping
	analyseGroup.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	analyseGroup.GET("/:id", handlers.GetAnalysePacienteHandler)
}
