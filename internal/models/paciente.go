package models

import "time"

type Paciente struct {
	ID                uint      `gorm:"column:paciente_id;primaryKey"`
	Nome              string    `gorm:"size:150;not null"`
	DataNascimento    time.Time `gorm:"type:date;not null"`
	Sexo              string    `gorm:"type:char(1);check:sexo IN ('M','F','O');not null"`
	Cpf               string    `gorm:"type:varchar(11);unique"`
	Telefone          string    `gorm:"type:varchar(20)"`
	Email             string    `gorm:"type:varchar(100)"`
	Endereco          string    `gorm:"type:text"`
	ContatoEmergencia string    `gorm:"column:contato_emergencia;size:100"`
	DataCadastro      time.Time `gorm:"autoCreateTime;column:data_cadastro"`

	Historicos []Historico `gorm:"foreignKey:PacienteID"`
	Exames     []Exame     `gorm:"foreignKey:PacienteID"`
	Consultas  []Consulta  `gorm:"foreignKey:PacienteID"`
}
