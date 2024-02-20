package db

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

func CreateRedisClient() (*redis.Client, error) {
	rdbHost, _ := base64.StdEncoding.DecodeString(os.Getenv("RDB_HOST"))
	rdbPort, _ := base64.StdEncoding.DecodeString(os.Getenv("RDB_PORT"))
	rdbPass := os.Getenv("RDB_PASS")
	rdbDbStr := os.Getenv("RDB_DB")

	urlRedis := fmt.Sprintf("redis://:%s@%s:%s/%s", rdbPass, rdbHost, rdbPort, rdbDbStr)
	opt, err := redis.ParseURL(urlRedis)
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)

	ctx := context.Background()
	_, err = client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	log.Println("[CreateRedisClient] Cliente Redis creado con éxito.")
	return client, nil
}
