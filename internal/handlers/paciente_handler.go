package handlers

import (
	"net/http"
	"saude/internal/models"
	"saude/internal/services"
	"strconv"

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

func UpdatePacienteHandler(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var paciente models.Paciente
	if err := ctx.ShouldBindJSON(&paciente); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	paciente.ID = uint(id)

	if err := services.UpdatePaciente(&paciente); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, paciente)
}

func DeletePacienteHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := services.DeletePaciente(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Paciente deletada com sucesso"})
}
