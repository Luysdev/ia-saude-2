package repository

import (
	"saude/internal/db"
	"saude/internal/models"
)

func CreateConsulta(c *models.Consulta) error {
	return db.DB.Create(c).Error
}

func GetAllConsulta() ([]models.Consulta, error) {
	var consulta []models.Consulta
	if err := db.DB.Find(&consulta).Error; err != nil {
		return nil, err
	}
	return consulta, nil
}

func GetConsultaById(id string) (*models.Consulta, error) {
	var consulta models.Consulta
	if err := db.DB.Where("consulta_id = ?", id).First(&consulta).Error; err != nil {
		return nil, err
	}
	return &consulta, nil
}
