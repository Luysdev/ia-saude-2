package services

import "saude/internal/models"

func CreateHistoricoService(h *models.Historico) error {
	return repository.CreateHistorico(h)
}
