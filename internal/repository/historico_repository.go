package repository

import (
	"saude/internal/db"
	"saude/internal/models"
)

func CreateHistorico(h *models.Historico) error {
	return db.DB.Create(h).Error
}

func GetAllHistorico() ([]models.Historico, error) {
	var historico []models.Historico
	if err := db.DB.Find(&historico).Error; err != nil {
		return nil, err
	}
	return historico, nil
}

func GetHistoricoById(id string) (*models.Historico, error) {
	var historico models.Historico
	if err := db.DB.Where("historico_id = ?", id).First(&historico).Error; err != nil {
		return nil, err
	}
	return &historico, nil
}

func UpdateHistorico(c *models.Historico) error {
	return db.DB.Updates(c).Error
}

func DeleteHistorico(id string) error {
	if err := db.DB.Where("historico_id = ?", id).Delete(&models.Historico{}).Error; err != nil {
		return err
	}
	return nil
}
