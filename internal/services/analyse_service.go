package services

import (
	"fmt"
	"saude/internal/ia"
	"saude/internal/models"
	"saude/internal/repository"
	"saude/internal/utils"
	"time"
)

func GetAnalysePacienteService(id string) (string, error) {
	paciente, err := repository.GetPacienteWithRelations(id)
	if err != nil {
		return "", err
	}

	return generatePacienteSummary(paciente)
}

func generatePacienteSummary(paciente *models.Paciente) (string, error) {

	basePrompt := utils.LoadPrompt("./internal/ia/prompts/base_summary.txt")
	pacientePrompt := utils.LoadPrompt("./internal/ia/prompts/paciente_summary.txt")

	pacienteData := fmt.Sprintf(
		"Nome: %s\nIdade: %d\nSexo: %s\nHistóricos: %v\nExames: %v\nConsultas: %v",
		paciente.Nome,
		int(time.Since(paciente.DataNascimento).Hours()/24/365),
		paciente.Sexo,
		paciente.Historicos,
		paciente.Exames,
		paciente.Consultas,
	)

	pacientePrompt = fmt.Sprintf(pacientePrompt, pacienteData)
	finalPrompt := basePrompt + "\n\n" + pacientePrompt

	resp, err := ia.AskGemini(finalPrompt)
	if err != nil {
		return "", err
	}

	return resp, nil
}
