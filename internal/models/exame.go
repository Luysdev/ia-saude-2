package models

import "time"

type Exame struct {
	ID               uint `gorm:"column:exame_id;primaryKey"`
	PacienteID       uint
	Paciente         Paciente  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Tipo             string    `gorm:"type:varchar(100)"`
	Data             time.Time `gorm:"type:date;not null"`
	Resultado        string    `gorm:"type:text"`
	ArquivoResultado []byte    `gorm:"column:arquivo_resultado"`
}
