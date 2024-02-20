package views

import "time"

type TramaCoban struct {
	IMEI           string
	Evento         string
	TimeStampEvent time.Time
	Latitud        float32
	Longitud       float32
	Velocidad      string
	Direccion      string
}
