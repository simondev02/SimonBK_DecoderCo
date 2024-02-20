package main

import (
	"SimonBK_DecoderCo/db"
	"SimonBK_DecoderCo/domain/service"
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("URL_RABBITMQ")

	if url == "" {
		panic("Error al leer las variables de entorno")
	}

	cfg := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := amqp.DialTLS(url, cfg)
	if err != nil {
		panic("Error al conectar a RabbitMQ: " + err.Error())
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic("Error al abrir un canal: " + err.Error())
	}
	defer ch.Close()

	queueName := "coban_queues"
	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic("Error al declarar la cola: " + err.Error())
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic("Error al registrar el consumidor: " + err.Error())
	}

	// Crear una conexión a Redis
	redisClient, err := db.CreateRedisClient()
	if err != nil {
		// Manejar el error
		log.Fatal("Error al conectar a Redis: ", err)
	}

	// Pasar la conexión a Redis a service.ParsearTrama
	go func() {
		for msg := range msgs {
			err := service.ParsearTrama(redisClient, string(msg.Body))
			if err != nil {
				fmt.Println("Error al parsear la trama:", err)
			} else {
				fmt.Println("Trama parseada:")
			}
		}
	}()

	log.Printf(" [*] Esperando mensajes. Para salir presione CTRL+C")
	select {}
}
