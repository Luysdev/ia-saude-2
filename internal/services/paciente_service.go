package services

import (
	"saude/internal/models"
	"saude/internal/repository"
)

func CreatePacienteService(h *models.Paciente) error {
	return repository.CreatePaciente(h)
}

func GetAllPacienteService() ([]models.Paciente, error) {
	return repository.GetAllPaciente()
}

func GetPacienteByIdService(id string) (*models.Paciente, error) {
	return repository.GetPacienteById(id)
}

func UpdatePaciente(c *models.Paciente) error {
	return repository.UpdatePaciente(c)
}

func DeletePaciente(id string) error {
	return repository.DeletePaciente(id)
}
