package handlers

import (
	"net/http"
	"saude/internal/models"
	"saude/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateExameHandler(c *gin.Context) {
	var exame models.Exame

	if err := c.ShouldBindJSON(&exame); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateExameService(&exame); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, exame)
}

func GetAllExameHandler(c *gin.Context) {
	exames, err := services.GetAllExameService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, exames)
}

func GetExameByIdHandler(c *gin.Context) {
	id := c.Param("id")

	exame, err := services.GetExameByIdService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, exame)
}

func UpdateExameHandler(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var exame models.Exame
	if err := ctx.ShouldBindJSON(&exame); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exame.ID = uint(id)

	if err := services.UpdateExame(&exame); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, exame)
}

func DeleteExameHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := services.DeleteExame(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Exame deletada com sucesso"})
}
