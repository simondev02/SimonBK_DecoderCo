package service

import (
	"SimonBK_DecoderCo/db"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v8"
)

func ObtenerCoordenadas(client *redis.Client, partes []string) (float64, float64) {

	if len(partes) < 11 {
		fmt.Println("No hay suficientes partes para obtener las coordenadas")
		return 0, 0
	}
	re := regexp.MustCompile("[^0-9]+")
	imei := re.ReplaceAllString(partes[0], "")

	// Verificar si las partes de la latitud y longitud están vacías
	if partes[7] == "" || partes[8] == "" || partes[9] == "" || partes[10] == "" {
		fmt.Println("Datos de latitud o longitud faltantes")
		latitud, longitud, err := ConsulLastLatLongInReddis(client, imei)
		if err != nil {
			fmt.Printf("Error al consultar el último registro en Redis: %v\n", err)
			fmt.Println("Intentando consultar en la base de datos...")
			latitud, longitud, err = ConsuLatLongInDb(db.DBConn, imei)
			if err != nil {
				fmt.Printf("Error al consultar el último registro en la base de datos: %v\n", err)
				return 0, 0
			}
		}
		return latitud, longitud
	}

	// Unir las partes de la latitud y longitud
	ubicacion := partes[7] + "," + partes[8] + "," + partes[9] + "," + partes[10]

	// Convertir las cadenas de latitud y longitud a float32
	latitud, longitud, err := convertirCoordenadas(ubicacion)
	if err != nil {
		fmt.Printf("Error al convertir las coordenadas: %v\n", err)
		return 0, 0
	}

	return latitud, longitud
}
func convertirCoordenadas(ubicacion string) (float64, float64, error) {
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
	fmt.Println("Latitud:", latitud, "Longitud:", longitud)
	return latitud, longitud, nil
}

func convertirCoordenada(coordStr string) (float64, error) {

	// Separar grados y minutos
	parts := strings.Split(coordStr, ".")
	grados, err := strconv.ParseFloat(parts[0][:len(parts[0])-2], 64)
	if err != nil {
		log.New(log.Writer(), "[convertirCoordenadas]", 0)
		return 0, nil
	}
	minutos, err := strconv.ParseFloat(parts[0][len(parts[0])-2:]+"."+parts[1], 64)
	if err != nil {
		return 0, nil
	}

	// Convertir a grados decimales
	coord := float64(grados + (minutos / 60.0))

	return coord, nil
}
