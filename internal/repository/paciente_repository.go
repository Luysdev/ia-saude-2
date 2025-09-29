package repository

import (
	"saude/internal/db"
	"saude/internal/models"
)

func CreatePaciente(p *models.Paciente) error {
	return db.DB.Create(p).Error
}

func GetAllPaciente() ([]models.Paciente, error) {
	var paciente []models.Paciente
	if err := db.DB.Find(&paciente).Error; err != nil {
		return nil, err
	}
	return paciente, nil
}

func GetPacienteById(id string) (*models.Paciente, error) {
	var paciente models.Paciente
	if err := db.DB.Where("Paciente_id = ?", id).First(&paciente).Error; err != nil {
		return nil, err
	}
	return &paciente, nil
}
