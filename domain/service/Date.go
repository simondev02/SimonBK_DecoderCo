package service

import (
	"fmt"
	"strings"
	"time"
)

func UnificarFechaHora(año string, mes string, dia string, hora string) (time.Time, error) {
	// Unificar año, mes, dia y hora en una sola cadena
	fechaHora := año + mes + dia + hora

	// Limpiar la cadena
	fechaHora = strings.TrimSpace(fechaHora)

	// Parsear la cadena en un valor de tiempo
	t, err := time.Parse("20060102150405", fechaHora)
	if err != nil {
		return time.Time{}, fmt.Errorf("error al parsear la fecha y la hora: %v", err)
	}

	return t, nil
}
