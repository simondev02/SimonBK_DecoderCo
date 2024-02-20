package service

import (
	"crypto/tls"
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

func DequeueFrame() error {

	// Obtener la URL de RabbitMQ de las variables de entorno
	url := os.Getenv("URL_RABBITMQ")

	if url == "" {
		return fmt.Errorf("error al leer las variables de entorno")
	}

	// Configurar TLS para saltar la verificación del certificado
	cfg := &tls.Config{
		InsecureSkipVerify: true,
	}

	// Conectarse al servidor RabbitMQ con TLS
	conn, err := amqp.DialTLS(url, cfg)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}
	defer conn.Close()

	// Crear un nuevo canal
	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %w", err)
	}
	defer ch.Close()

	// Declarar la cola
	queueName := "coban_queues"
	q, err := ch.QueueDeclare(
		queueName, // name of the queue
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %w", err)
	}

	// Consumir mensajes de la cola
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %w", err)
	}

	// Obtener un solo mensaje de la cola
	msg := <-msgs

	// Imprimir el mensaje en la consola
	fmt.Println(string(msg.Body))

	return nil
}
