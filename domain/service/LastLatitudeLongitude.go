package service

import (
	views "SimonBK_DecoderCo/api/views"
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func ConsulLastLatLongInReddis(client *redis.Client, imei string) (float64, float64, error) {
	ctx := context.Background()

	// Consultar el último registro en el stream 'avl_service'
	result, err := client.XRevRangeN(ctx, "avl_service", "+", "-", 1).Result()
	if err != nil {
		return 0, 0, fmt.Errorf("error al consultar el último registro: %v", err)
	}

	// Buscar el registro con el IMEI especificado
	for _, message := range result {
		var record views.AvlRecord
		err := json.Unmarshal([]byte(message.Values["event"].(string)), &record)
		if err != nil {
			return 0, 0, fmt.Errorf("error al deserializar el registro: %v", err)
		}

		if record.Imei == imei {
			return record.Latitude, record.Longitude, nil
		}
	}

	// Si no se encontró un registro con el IMEI especificado, retorna 0 para la latitud y la longitud
	return 0, 0, nil
}
