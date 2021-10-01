package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type Conn struct {
	Channel *amqp.Channel
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s", msg, err)
	}
}

func GetConn(amqp_url string) (*Conn, error) {
	conn, err := amqp.Dial(amqp_url)
	failOnError(err, "error: Dials connect error")

	ch, err := conn.Channel()
	failOnError(err, "error: create channel")

	return &Conn{
		Channel: ch,
	}, nil
}
