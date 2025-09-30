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

	pacienteGroup.POST("/", handlers.CreatePacienteHandler)
	pacienteGroup.GET("/", handlers.GetAllPacienteHandler)
	pacienteGroup.GET("/:id", handlers.GetPacienteByIdHandler)
	pacienteGroup.PUT("/:id", handlers.UpdatePacienteHandler)
	pacienteGroup.DELETE("/:id", handlers.DeletePacienteHandler)
}
