package db

import (
	"log"
	"saude/internal/models"
)

func Migrate() {
	err := DB.AutoMigrate(
		&models.Paciente{},
		&models.Medico{},
		&models.Historico{},
		&models.Exame{},
	)

	if err != nil {
		log.Fatal("❌ Erro na migração das tabelas:", err)
	}

	log.Println("✅ Tabelas criadas/atualizadas com sucesso!")
}
