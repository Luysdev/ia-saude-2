package handlers

import (
	"net/http"
	"saude/internal/models"
	"saude/internal/services"
	"strconv"

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

func UpdateMedicoHandler(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var medico models.Medico
	if err := ctx.ShouldBindJSON(&medico); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	medico.ID = uint(id)

	if err := services.UpdateMedico(&medico); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, medico)
}

func DeleteMedicoHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := services.DeleteMedico(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Medico deletada com sucesso"})
}
