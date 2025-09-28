package services

import (
	"saude/internal/models"
	"saude/internal/repository"
)

func CreateConsultaService(h *models.Consulta) error {
	return repository.CreateConsulta(h)
}

func GetAllConsultaService() ([]models.Consulta, error) {
	return repository.GetAllConsulta()
}

func GetConsultaByIdService(id string) (*models.Consulta, error) {
	return repository.GetConsultaById(id)
}
