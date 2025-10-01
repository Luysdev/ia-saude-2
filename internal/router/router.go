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
	routes.RoutePaciente(router)
	routes.RouteConsulta(router)
	routes.RouteAnalyse(router)
	routes.RouteExame(router)

	router.Run(":8080")
}
