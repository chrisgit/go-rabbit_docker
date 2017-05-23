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

func sendMessage(messageBody string) {
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

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(messageBody),
		})
	log.Printf(" [x] Sent %s", messageBody)
	failOnError(err, "Failed to publish a message")
}
