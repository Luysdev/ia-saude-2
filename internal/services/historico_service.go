package services

import (
	"saude/internal/models"
	"saude/internal/repository"
)

func CreateHistoricoService(h *models.Historico) error {
	return repository.CreateHistorico(h)
}

func GetAllHistoricoService() ([]models.Historico, error) {
	return repository.GetAllHistorico()
}

func GetHistoricoByIdService(id string) (*models.Historico, error) {
	return repository.GetHistoricoById(id)
}
