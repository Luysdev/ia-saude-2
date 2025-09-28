package models

import "time"

type Medico struct {
	ID            uint      `gorm:"column:medico_id;primaryKey"`
	Nome          string    `gorm:"size:150;not null"`
	Crm           string    `gorm:"size:20;unique;not null"`
	Especialidade string    `gorm:"size:100;not null"` // Pode virar ENUM se quiser
	Telefone      string    `gorm:"size:20"`
	Email         string    `gorm:"size:100"`
	Endereco      string    `gorm:"type:text"`
	DataCadastro  time.Time `gorm:"autoCreateTime;column:data_cadastro"`
}
