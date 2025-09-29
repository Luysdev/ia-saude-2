package ia

import (
	"context"
	"fmt"
	"log"
	"os"

	genai "github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var client *genai.Client

func InitGemini() {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("⚠️ Defina a variável de ambiente GEMINI_API_KEY com sua chave")
	}

	ctx := context.Background()
	var err error
	client, err = genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Erro criando cliente Gemini: %v", err)
	}
}

// AskGemini envia um prompt e retorna a resposta
func AskGemini(prompt string) (string, error) {
	ctx := context.Background()
	model := client.GenerativeModel("models/gemini-2.5-flash")

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))

	if err != nil {
		return "", err
	}

	var result string
	for _, candidate := range resp.Candidates {
		if candidate.Content != nil {
			for _, part := range candidate.Content.Parts {
				result += fmt.Sprintf("%v", part)
			}
		}
	}
	return result, nil
}
