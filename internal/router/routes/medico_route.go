package routes

import (
	"saude/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RouteMedico(router *gin.Engine) {

	medicoGroup := router.Group("/medico")

	// Endpoint de ping
	medicoGroup.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	medicoGroup.POST("/", handlers.CreateMedicoHandler)
	medicoGroup.GET("/", handlers.GetAllMedicosHandler)
	medicoGroup.GET("/:id", handlers.GetMedicoByIdHandler)
	medicoGroup.PUT("/:id", handlers.UpdateMedicoHandler)
	medicoGroup.DELETE("/:id", handlers.DeleteMedicoHandler)

}
