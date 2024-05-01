package main

import (
	"docx-doc-pipeline-srv/src/pkg/rabbit"
	"log"
	"log/slog"
)

func main() {
	amqp := rabbit.NewRabbit("amqp://guest:guest@localhost:5672", int64(10000))
	if err := amqp.Connect(); err != nil {
		slog.Error(err.Error())
	}
	amqpChannel := amqp.GetChannel()
	if err := amqpChannel.ExchangeDeclare("pipeline", "topic", true, false, false, false, nil); err != nil {
		slog.Error(err.Error())
	}

	q, _ := amqpChannel.QueueDeclare(
		"pdf.merge.request.queue",
		false,
		false,
		true,
		false,
		nil,
	)

	amqpChannel.QueueBind(
		q.Name,
		"pdf.merge.request.*",
		"pipeline",
		false,
		nil,
	)

	msgs, _ := amqpChannel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
