package repository

import (
	"saude/internal/db"
	"saude/internal/models"
)

func CreateExame(e *models.Exame) error {
	return db.DB.Create(e).Error
}

func GetAllExame() ([]models.Exame, error) {
	var exame []models.Exame
	if err := db.DB.Find(&exame).Error; err != nil {
		return nil, err
	}
	return exame, nil
}

func GetExameById(id string) (*models.Exame, error) {
	var exame models.Exame
	if err := db.DB.Where("exame_id = ?", id).First(&exame).Error; err != nil {
		return nil, err
	}
	return &exame, nil
}

func UpdateExame(e *models.Exame) error {
	return db.DB.Updates(e).Error
}

func DeleteExame(id string) error {
	if err := db.DB.Where("exame_id = ?", id).Delete(&models.Exame{}).Error; err != nil {
		return err
	}
	return nil
}
