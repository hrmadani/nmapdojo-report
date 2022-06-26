package controllers

import (
	"log"

	"github.com/streadway/amqp"
)

//Constants
const (
	RabbitMQName             = "user_report"
	RabbitMQDurable          = false
	RabbitMQDeleteWhenUnused = false
	RabbitMQExclusive        = false
	RabbitMQNoWait           = false
	RabbitMQConsumer         = ""
	RabbitMQAutoAck          = true
	RabbitMQNoLocal          = false
)

//Error Handler
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//The main function in this service
//Consume messages from RabbitMQ
//Call the appropriate function to save messages to database
func ConsumeFromRabbit() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		RabbitMQName,
		RabbitMQDurable,
		RabbitMQDeleteWhenUnused,
		RabbitMQExclusive,
		RabbitMQNoWait,
		nil, // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,
		RabbitMQConsumer,
		RabbitMQAutoAck,
		RabbitMQExclusive,
		RabbitMQNoLocal,
		RabbitMQNoWait,
		nil, // args
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
