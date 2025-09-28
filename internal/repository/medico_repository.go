package repository

import (
	"saude/internal/db"
	"saude/internal/models"
)

func CreateMedico(m *models.Medico) error {
	return db.DB.Create(m).Error
}

func GetAllMedicos() ([]models.Medico, error) {
	var medicos []models.Medico
	if err := db.DB.Find(&medicos).Error; err != nil {
		return nil, err
	}
	return medicos, nil
}

func GetMedicoById(id string) (*models.Medico, error) {
	var medico models.Medico
	if err := db.DB.Where("medico_id = ?", id).First(&medico).Error; err != nil {
		return nil, err
	}

	return &medico, nil
}
