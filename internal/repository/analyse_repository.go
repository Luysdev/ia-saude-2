package repository

import (
	"saude/internal/db"
	"saude/internal/models"
)

func GetPacienteWithRelations(id string) (*models.Paciente, error) {
	var paciente models.Paciente

	err := db.DB.Preload("Historicos").
		Preload("Consultas").
		Preload("Exames").
		First(&paciente, "paciente_id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &paciente, nil
}
