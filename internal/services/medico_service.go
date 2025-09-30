package services

import (
	"saude/internal/models"
	"saude/internal/repository"
)

// Service responsável por criar médico
func CreateMedicoService(m *models.Medico) error {
	return repository.CreateMedico(m)
}

func GetAllMedicosService() ([]models.Medico, error) {
	return repository.GetAllMedicos()
}

func GetMedicoByIdService(id string) (*models.Medico, error) {
	return repository.GetMedicoById(id)
}

func UpdateMedico(c *models.Medico) error {
	return repository.UpdateMedico(c)
}

func DeleteMedico(id string) error {
	return repository.DeleteMedico(id)
}
