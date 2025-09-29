package routes

import (
	"saude/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RoutePaciente(router *gin.Engine) {
	pacienteGroup := router.Group("/paciente")

	// Endpoint de ping
	pacienteGroup.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// // *CREATE*
	pacienteGroup.POST("/", handlers.CreatePacienteHandler)
	// *GETALL*
	pacienteGroup.GET("/", handlers.GetAllPacienteHandler)
	// *GETBYID*
	pacienteGroup.GET("/:id", handlers.GetPacienteByIdHandler)
}
