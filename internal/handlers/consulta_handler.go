package handlers

import (
	"net/http"
	"saude/internal/models"
	"saude/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CREATE
func CreateConsultaHandler(ctx *gin.Context) {
	var consulta models.Consulta
	if err := ctx.ShouldBindJSON(&consulta); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateConsulta(&consulta); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, consulta)
}

// GET ALL
func GetAllConsultaHandler(ctx *gin.Context) {
	consultas, err := services.GetAllConsulta()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, consultas)
}

// GET BY ID
func GetConsultaByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	consulta, err := services.GetConsultaById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Consulta não encontrada"})
		return
	}
	ctx.JSON(http.StatusOK, consulta)
}

func UpdateConsultaHandler(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var consulta models.Consulta
	if err := ctx.ShouldBindJSON(&consulta); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	consulta.ID = uint(id)

	if err := services.UpdateConsulta(&consulta); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, consulta)
}

func DeleteConsultaHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := services.DeleteConsulta(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Consulta deletada com sucesso"})
}
