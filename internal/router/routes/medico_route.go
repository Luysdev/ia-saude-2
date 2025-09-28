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

	// *CREATE*
	medicoGroup.POST("/", handlers.CreateMedicoHandler)
	// *GETALL*
	medicoGroup.GET("/", handlers.GetAllMedicosHandler)
	// *GETBYID*
	medicoGroup.GET("/:id", handlers.GetMedicoByIdHandler)
	// // *UPDATE*
	// medicoGroup.PUT("/")
	// // *DELETE*
	// medicoGroup.DELETE("/")

}
