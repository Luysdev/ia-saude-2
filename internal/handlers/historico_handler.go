package handlers

import (
	"net/http"
	"saude/internal/models"
	"saude/internal/services"

	"github.com/gin-gonic/gin"
)

func CreateHistoricoHandler(c *gin.Context) {
	var historico models.Historico

	if err := c.ShouldBindJSON(&historico); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateHistoricoService(&historico); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, historico)
}

func GetAllHistoricoHandler(c *gin.Context) {
	historicos, err := services.GetAllHistoricoService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, historicos)
}

func GetHistoricoByIdHandler(c *gin.Context) {
	id := c.Param("id")

	historico, err := services.GetHistoricoByIdService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, historico)
}
