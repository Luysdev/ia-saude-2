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

func UpdateHistorico(c *models.Historico) error {
	return repository.UpdateHistorico(c)
}

func DeleteHistorico(id string) error {
	return repository.DeleteHistorico(id)
}
