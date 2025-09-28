package models

import "time"

type Historico struct {
	ID              uint `gorm:"column:historico_id;primaryKey"`
	PacienteID      uint
	Paciente        Paciente  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Condicao        string    `gorm:"type:varchar(200)"`
	Descricao       string    `gorm:"type:text"`
	DataDiagnostico time.Time `gorm:"type:date"`
	Status          bool      `gorm:"default:true"`
}
