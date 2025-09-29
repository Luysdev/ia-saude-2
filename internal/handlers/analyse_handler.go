package handlers

import (
	"net/http"
	"saude/internal/services"

	"github.com/gin-gonic/gin"
)

func GetAnalysePacienteHandler(c *gin.Context) {
	id := c.Param("id")

	data, err := services.GetAnalysePacienteService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"info": data,
	})
}
