package handlers

import (
	"net/http"
	"saude/internal/models"
	"saude/internal/services"

	"github.com/gin-gonic/gin"
)

func CreatePacienteHandler(c *gin.Context) {
	var paciente models.Paciente

	if err := c.ShouldBindJSON(&paciente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreatePacienteService(&paciente); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, paciente)
}

func GetAllPacienteHandler(c *gin.Context) {
	pacientes, err := services.GetAllPacienteService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pacientes)
}

func GetPacienteByIdHandler(c *gin.Context) {
	id := c.Param("id")

	paciente, err := services.GetPacienteByIdService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, paciente)
}
