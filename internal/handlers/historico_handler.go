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
