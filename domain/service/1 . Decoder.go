package service

import (
	views "SimonBK_DecoderCo/api/views"
	"SimonBK_DecoderCo/db"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

func ParsearTrama(client *redis.Client, trama string) error {
	println(trama)
	partes := strings.Split(trama, ",")

	// 1 . Imei
	imei := strings.TrimPrefix(partes[0], "imei:")

	// 2 . Evento
	evento, err := GetEvent(partes[1])
	if err != nil {
		fmt.Printf("Error al obtener el evento: %v\n", err)
		evento = ""
	}

	// 3 . Fecha Hora Evento
	fechaHora, err := ParsearFechaHora(partes[2])
	if err != nil {
		fmt.Printf("Error al parsear fecha y hora: %v\n", err)
		fechaHora = time.Time{} // Fecha y hora vacía
	}

	// 4 . Coordenadas (Latitud, Longitud)
	latitud, longitud := ObtenerCoordenadas(client, partes)

	// 5 .Velocidad
	velocidad, err := GetSpeed(partes[11])
	if err != nil {
		fmt.Printf("Error al obtener la velocidad: %v\n", err)
		velocidad = 0
	}

	// 6 . Obtener la información del vehículo
	info, err := GetVehiclesInfo(imei)
	if err != nil {
		fmt.Printf("Error al obtener la información del vehículo: %v\n", err)
		info = views.VehicleInfo{} // Información del vehículo vacía
	}

	// 7 . Angulo
	angulo, err := ObtenerAngulo(partes[12])
	if err != nil {
		fmt.Printf("Error al obtener el ángulo: %v\n", err)
		angulo = 0
	}
	// 8 . Validez
	//validez := partes[10]

	// El resto de tu código va aquí...

	// Crear la estructura TramaCoban con los datos obtenidos
	tramaCoban := views.AvlRecord{
		Plate:          info.Plate,
		Imei:           imei,
		Event:          evento,
		Id_company:     info.IdCompany,
		Company:        info.Company,
		Id_customer:    info.IdCustomer,
		Customer:       info.Customer,
		TimeStampEvent: fechaHora,
		Latitude:       latitud,
		Longitude:      longitud,
		Angle:          angulo,
		Speed:          velocidad,
	}

	// Manejamos el estado de vehiculo encendido o vehiculo apagado.
	var ignition int
	// Verificar el valor de 'Event'
	if evento == "Vehiculo encendido" {
		ignition = 1
	} else if evento == "Vehiculo apagado" {
		ignition = 0
	} else {
		// Consultar el último registro
		var err error
		ignition, err = ConsultarUltimoRegistro(client, imei)
		if err != nil {
			log.Fatalf("Error al consultar el último registro: %v", err)
		}
	}
	// Asignar el valor de Ignition a AnalogInput1
	properties := views.Properties{
		Ignition: ignition,
		Speed:    int(velocidad),
	}

	propertiesJSON, err := json.Marshal(properties)
	if err != nil {
		return fmt.Errorf("error al convertir Properties a JSON: %v", err)
	}

	tramaCoban.Properties = string(propertiesJSON)

	// Insertamos la trama en Redis
	redisClient, err := db.CreateRedisClient()
	if err != nil {
		log.Fatal("Error al conectar a Redis: ", err)
	}
	err = InsertarEnRedis(redisClient, tramaCoban)
	if err != nil {
		return fmt.Errorf("error al insertar en Redis: %v", err)
	}

	return nil
}
