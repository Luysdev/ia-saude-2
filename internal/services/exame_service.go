package services

import (
	"saude/internal/models"
	"saude/internal/repository"
)

func CreateExameService(h *models.Exame) error {
	return repository.CreateExame(h)
}

func GetAllExameService() ([]models.Exame, error) {
	return repository.GetAllExame()
}

func GetExameByIdService(id string) (*models.Exame, error) {
	return repository.GetExameById(id)
}
