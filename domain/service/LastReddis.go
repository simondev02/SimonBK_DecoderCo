package service

import (
	"context"
	"encoding/json"
	"fmt"

	views "SimonBK_DecoderCo/api/views"

	"github.com/go-redis/redis/v8"
)

func ConsultarUltimoRegistro(client *redis.Client, imei string) (int, error) {
	ctx := context.Background()

	// Consultar el último registro en el stream 'avl_service'
	result, err := client.XRevRangeN(ctx, "avl_service", "+", "-", 1).Result()
	if err != nil {
		return 0, fmt.Errorf("error al consultar el último registro: %v", err)
	}

	// Buscar el registro con el IMEI especificado
	for _, message := range result {
		var record views.AvlRecord
		err := json.Unmarshal([]byte(message.Values["event"].(string)), &record)
		if err != nil {
			return 0, fmt.Errorf("error al deserializar el registro: %v", err)
		}

		if record.Imei == imei {
			var properties views.Properties
			err := json.Unmarshal([]byte(record.Properties), &properties)
			if err != nil {
				return 0, fmt.Errorf("error al deserializar Properties: %v", err)
			}

			return properties.Ignition, nil
		}
	}

	// Si no se encontró un registro con el IMEI especificado, retorna 0
	return 0, nil
}
