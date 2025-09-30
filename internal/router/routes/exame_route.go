package routes

import (
	"saude/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RouteExame(router *gin.Engine) {
	exameGroup := router.Group("/exame")

	exameGroup.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	exameGroup.POST("/", handlers.CreateExameHandler)
	exameGroup.GET("/", handlers.GetAllExameHandler)
	exameGroup.GET("/:id", handlers.GetExameByIdHandler)
	exameGroup.PUT("/:id", handlers.UpdateExameHandler)
	exameGroup.DELETE("/:id", handlers.DeleteExameHandler)
}
