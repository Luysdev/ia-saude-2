package models

import "time"

type Exame struct {
	ID               uint      `gorm:"column:exame_id;primaryKey" json:"id"`
	PacienteID       uint      `json:"paciente_id"`
	Paciente         Paciente  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Tipo             string    `gorm:"type:varchar(100)" json:"tipo"`
	Data             time.Time `gorm:"type:date;not null" json:"data"`
	Resultado        string    `gorm:"type:text" json:"resultado"`
	ArquivoResultado []byte    `gorm:"column:arquivo_resultado" json:"arquivo_resultado"`
}
