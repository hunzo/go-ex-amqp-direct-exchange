package main

import "go-ex-amqp-direct-exchange/rabbitmq"

func main() {
	con, err := rabbitmq.GetConn("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err)
	}
	// con.PublishDirect()
	con.ConsumeDirect()
}
