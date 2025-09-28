package config

import (
	"log"
	"saude/internal/db"

	"github.com/joho/godotenv"
)

func InitConfig() {

	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env não carregado, usando variáveis do sistema")
	}

	db.Connect()
	db.Migrate()
}
