package services

import (
	"saude/internal/models"
	"saude/internal/repository"
)

func CreateConsulta(c *models.Consulta) error {
	return repository.CreateConsulta(c)
}

func GetAllConsulta() ([]models.Consulta, error) {
	return repository.GetAllConsulta()
}

func GetConsultaById(id string) (*models.Consulta, error) {
	return repository.GetConsultaById(id)
}

func UpdateConsulta(c *models.Consulta) error {
	return repository.UpdateConsulta(c)
}

func DeleteConsulta(id string) error {
	return repository.DeleteConsulta(id)
}
