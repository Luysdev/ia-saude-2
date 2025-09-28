package router

import (
	"saude/internal/router/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	router.Use(cors.Default())

	routes.RouteMedico(router)
	routes.RouteHistorico(router)

	router.Run(":4949")
}
