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

func UpdateExame(e *models.Exame) error {
	return repository.UpdateExame(e)
}

func DeleteExame(id string) error {
	return repository.DeleteExame(id)
}
