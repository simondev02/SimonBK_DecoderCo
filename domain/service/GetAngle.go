package service

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ObtenerAngulo(anguloStr string) (int16, error) {
	anguloStr = strings.TrimPrefix(anguloStr, "imei:")

	// Limpiar la cadena
	re := regexp.MustCompile("[^0-9.]+")
	anguloStr = re.ReplaceAllString(anguloStr, "")

	if anguloStr == "" {
		return 0, fmt.Errorf("dato de ángulo faltante")
	}

	anguloFloat, err := strconv.ParseFloat(anguloStr, 64)
	if err != nil {
		return 0, fmt.Errorf("error al convertir el ángulo a float: %v", err)
	}

	return int16(anguloFloat), nil
}
