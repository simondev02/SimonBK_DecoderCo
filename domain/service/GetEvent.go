package service

import (
	"SimonBK_DecoderCo/domain/models"
)

func GetEvent(registro string) (string, error) {
	evento, existe := models.EventosCoban[registro]
	if !existe {
		// Retorna el mismo registro si no se encuentra un evento
		return registro, nil
	}
	return evento, nil
}
