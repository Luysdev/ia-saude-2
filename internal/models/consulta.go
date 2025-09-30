package models

import "time"

type Consulta struct {
	ID            uint      `gorm:"column:consulta_id;primaryKey" json:"id"`
	PacienteID    uint      `gorm:"column:paciente_id" json:"paciente_id"`
	Paciente      Paciente  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"paciente,omitempty"`
	Especialidade string    `gorm:"type:varchar(100)" json:"especialidade"`
	DataConsulta  time.Time `gorm:"type:timestamp;not null" json:"data_consulta"`
	Sintomas      string    `gorm:"type:text" json:"sintomas"`
	Diagnostico   string    `gorm:"type:text" json:"diagnostico"`
	Prescricao    string    `gorm:"type:text" json:"prescricao"`
}
