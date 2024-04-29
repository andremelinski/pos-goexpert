package topicexchange

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel()(*amqp.Channel, error){
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		panic(err)
	}

	rabbitmqChan, err := conn.Channel()

	if err != nil {
		panic(err)
	}

	return rabbitmqChan, err
}

func Publisher(chn *amqp.Channel, body, exchangeName, key string )(error){
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := chn.PublishWithContext(
		ctx,
		exchangeName,
		key,
		false,
		false,
	amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

	if err != nil {
		return err
	}

	return nil
}

// func Consumer()(string, error){

// }

