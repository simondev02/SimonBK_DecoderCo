package service

import (
	views "SimonBK_DecoderCo/api/views"
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func InsertarEnRedis(client *redis.Client, record views.AvlRecord) error {
	// Convertir la estructura AvlRecord a JSON
	jsonRecord, err := json.Marshal(record)
	if err != nil {
		return fmt.Errorf("error al convertir AvlRecord a JSON: %v", err)
	}

	// Insertar el JSON en el canal de Redis
	_, err = client.XAdd(ctx, &redis.XAddArgs{
		Stream: "avl_service_stream",
		Values: map[string]interface{}{"event": string(jsonRecord)},
	}).Result()
	if err != nil {
		return fmt.Errorf("error al insertar en Redis: %v", err)
	}

	// Publicar el JSON en el canal de Redis 'avl_service'
	_, err = client.Publish(ctx, "avl_service", string(jsonRecord)).Result()
	if err != nil {
		return fmt.Errorf("error al publicar en Redis: %v", err)
	}

	return nil
}
