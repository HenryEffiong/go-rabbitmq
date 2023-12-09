package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqps://student:XYR4yqc.cxh4zug6vje@rabbitmq-exam.rmq3.cloudamqp.com/mxifnklj")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// create a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// err = ch.ExchangeDeclare(
	// 	"exchange.8811fb07-d8ad-4c36-b3b2-0a9c71411d5f", // name
	// 	"direct", // type
	// 	false,    // durable
	// 	true,     // auto-deleted
	// 	false,    // internal
	// 	false,    // no-wait
	// 	nil,      // arguments
	// )
	// failOnError(err, "Failed to declare an exchange")

	err = ch.QueueBind(
			"exam",        // queue name
			"8811fb07-d8ad-4c36-b3b2-0a9c71411d5f",             // routing key
			"exchange.8811fb07-d8ad-4c36-b3b2-0a9c71411d5f", // exchange
			false,
			nil)
		failOnError(err, "Failed to bind a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"exchange.8811fb07-d8ad-4c36-b3b2-0a9c71411d5f", // exchange
		"8811fb07-d8ad-4c36-b3b2-0a9c71411d5f",          // routing key
		false,                                           // mandatory
		false,                                           // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte("Hi CloudAMQP, this was fun!"),
		})
	failOnError(err, "Failed to publish a message")
	log.Println("Done")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
