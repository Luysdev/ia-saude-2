package utils

import (
	"io"
	"log"
	"os"
)

func LoadPrompt(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Erro ao abrir prompt %s: %v", filename, err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Erro ao ler prompt %s: %v", filename, err)
	}

	return string(data)
}
