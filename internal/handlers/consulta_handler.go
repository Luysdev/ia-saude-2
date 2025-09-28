package handlers

import (
	"net/http"
	"saude/internal/models"
	"saude/internal/services"

	"github.com/gin-gonic/gin"
)

func CreateConsultaHandler(c *gin.Context) {
	var consulta models.Consulta

	if err := c.ShouldBindJSON(&consulta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateConsultaService(&consulta); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, consulta)
}

func GetAllConsultaHandler(c *gin.Context) {
	consultas, err := services.GetAllConsultaService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, consultas)
}

func GetConsultaByIdHandler(c *gin.Context) {
	id := c.Param("id")

	consulta, err := services.GetConsultaByIdService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, consulta)
}
