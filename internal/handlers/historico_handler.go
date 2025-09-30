package handlers

import (
	"net/http"
	"saude/internal/models"
	"saude/internal/services"
	"strconv"

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

func UpdateHistoricoHandler(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var historico models.Historico
	if err := ctx.ShouldBindJSON(&historico); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	historico.ID = uint(id)

	if err := services.UpdateHistorico(&historico); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, historico)
}

func DeleteHistoricoHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := services.DeleteHistorico(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Historico deletada com sucesso"})
}
