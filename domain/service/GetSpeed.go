package service

import (
	"fmt"
	"strconv"
)

func GetSpeed(velocidadStr string) (int16, error) {
	if velocidadStr == "" {
		return 0, fmt.Errorf("dato de velocidad faltante")
	}

	velocidadFloat, err := strconv.ParseFloat(velocidadStr, 64)
	if err != nil {
		return 0, fmt.Errorf("error al convertir la velocidad a float: %v", err)
	}

	return int16(velocidadFloat), nil
}
