package service

import (
	"fmt"
	"strconv"
	"strings"
)

func ObtenerCoordenadas(partes []string) (float32, float32, error) {
	if len(partes) < 11 {
		return 0, 0, fmt.Errorf("no hay suficientes partes para obtener las coordenadas")
	}

	// Unir las partes de la latitud y longitud
	ubicacion := partes[7] + "," + partes[8] + "," + partes[9] + "," + partes[10]

	if ubicacion == "" {
		return 0, 0, fmt.Errorf("datos de latitud o longitud faltantes")
	}

	// Convertir las cadenas de latitud y longitud a float32
	latitud, longitud, err := convertirCoordenadas(ubicacion)
	if err != nil {
		return 0, 0, fmt.Errorf("error al convertir las coordenadas: %v", err)
	}

	return latitud, longitud, nil
}
func convertirCoordenadas(ubicacion string) (float32, float32, error) {
	// Separar la ubicación en latitud y longitud
	partes := strings.Split(ubicacion, ",")
	if len(partes) < 4 {
		return 0, 0, fmt.Errorf("ubicación inválida")
	}

	latitudStr := partes[0]
	longitudStr := partes[2]

	// Convertir las cadenas de latitud y longitud a float32
	latitud, err := convertirCoordenada(latitudStr)
	if err != nil {
		return 0, 0, fmt.Errorf("error al convertir la latitud: %v", err)
	}

	longitud, err := convertirCoordenada(longitudStr)
	if err != nil {
		return 0, 0, fmt.Errorf("error al convertir la longitud: %v", err)
	}

	// Ajustar la dirección
	if partes[1] == "S" {
		latitud = -latitud
	}
	if partes[3] == "W" {
		longitud = -longitud
	}

	return latitud, longitud, nil
}

func convertirCoordenada(coordStr string) (float32, error) {
	// Separar grados y minutos
	parts := strings.Split(coordStr, ".")
	grados, err := strconv.ParseFloat(parts[0][:len(parts[0])-2], 64)
	if err != nil {
		return 0, err
	}
	minutos, err := strconv.ParseFloat(parts[0][len(parts[0])-2:]+"."+parts[1], 64)
	if err != nil {
		return 0, err
	}

	// Convertir a grados decimales
	coord := float32(grados + (minutos / 60.0))

	return coord, nil
}
