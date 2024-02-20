package service

import (
	"fmt"
	"time"
)

func ParsearFechaHora(fechaHoraParte string) (time.Time, error) {
	// Parsear la fecha y hora del evento.
	fechaHora, err := time.Parse("060102150405", fechaHoraParte)
	if err != nil {
		return time.Time{}, fmt.Errorf("[Decoder] error al parsear fecha y hora: %v", err)
	}

	return fechaHora, nil
}
