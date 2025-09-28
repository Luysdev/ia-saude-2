package models

import "time"

type Consulta struct {
	ID            uint `gorm:"column:consulta_id;primaryKey"`
	PacienteID    uint
	Paciente      Paciente  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Especialidade string    `gorm:"type:varchar(100)"`
	DataConsulta  time.Time `gorm:"type:timestamp;not null"`
	Sintomas      string    `gorm:"type:text"`
	Diagnostico   string    `gorm:"type:text"`
	Prescricao    string    `gorm:"type:text"`
}
