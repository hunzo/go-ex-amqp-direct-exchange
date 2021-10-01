package rabbitmq

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func (conn Conn) PublishDirect() {

	q, err := conn.Channel.QueueDeclare(
		"queue-name",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "error: queue declare")

	body := `{
		"message": "hello world",
		"timestamp": ` + time.Now().String() + `
	}`

	err = conn.Channel.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         []byte(body),
			DeliveryMode: amqp.Persistent,
		},
	)
	failOnError(err, "failed on Publish")

	fmt.Printf("Message send: %s\n", body)
}

func (conn Conn) ConsumeDirect() {
	q, err := conn.Channel.QueueDeclare(
		"queue-name",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "failed to declare queue")

	msqs, err := conn.Channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "failed to register consumer")

	forever := make(chan bool)

	go func() {
		for d := range msqs {
			log.Printf("Recived a message: %s", d.Body)
		}
	}()

	log.Printf("Start Comsumer !!")

	<-forever

}
