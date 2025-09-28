package handlers

import (
	"net/http"
	"saude/internal/models"
	"saude/internal/services"

	"github.com/gin-gonic/gin"
)

func CreateMedicoHandler(c *gin.Context) {
	var medico models.Medico

	if err := c.ShouldBindJSON(&medico); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateMedicoService(&medico); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, medico)
}

func GetAllMedicosHandler(c *gin.Context) {
	medicos, err := services.GetAllMedicosService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, medicos)
}

func GetMedicoByIdHandler(c *gin.Context) {
	id := c.Param("id")

	medicos, err := services.GetMedicoByIdService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, medicos)
}
