package models

var EventosCoban = map[string]string{
	"001":             "Información de ubicación", // "Location information"
	"help me":         "Alarma SOS",               // "SOS alarm"
	"low battery":     "Batería baja",             // "Low battery alarm"
	"move":            "Vehiculo en Movimiento",   // "Movement alarm"
	"speed":           "Exceso de velocidad",      // "Over speed alarm"
	"stockade":        "GEO-cerca",                // "GEO-fence alarm"
	"ac alarm":        "Apagado",
	"acc on":          "Vehiculo encendido",                         // "Power off alarm"
	"acc off":         "Vehiculo apagado",                           // "Power off alarm"
	"door alarm":      "Puerta abierta",                             // "Door open alarm"
	"sensor alarm":    "Choque",                                     // "Shock alarm"
	"acc alarm":       "Encendido/Apagado",                          // "Acc alarm"
	"accident alarm":  "Accidente",                                  // "Accident alarm"
	"bonnet alarm":    "Capó",                                       // "Bonnet alarm"
	"footbrake alarm": "Freno de pie",                               // "Footbrake alarm"
	"T":               "Temperatura",                                // "Temperature alarm"
	"oil":             "Combustible",                                // "Fuel alarm"
	"DTC P0001":       "Código de problema de diagnóstico: P0001",   // "Diagnostic trouble code: P0001"
	"service":         "Notificación de mantenimiento del vehículo", // "Vehicle maintenance notification"
}
