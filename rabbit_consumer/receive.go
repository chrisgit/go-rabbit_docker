package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func rabbitHostname() string {
	rabbitHostname := os.Getenv("RABBIT_HOSTNAME")
	if rabbitHostname != "" {
		return rabbitHostname
	}
	return *rabbitHostPtr
}

func rabbitPort() int {
	rabbitPort := os.Getenv("RABBIT_PORT")
	if rabbitPort != "" {
		rabbitPort, _ := strconv.Atoi(rabbitPort)
		return rabbitPort
	}
	return *rabbitPortPtr
}

func rabbitAMQP() string {
	return fmt.Sprintf("amqp://guest:guest@%s:%d/", rabbitHostname(), rabbitPort())
}

func receiveMessage() {
	conn, err := amqp.Dial(rabbitAMQP())
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
