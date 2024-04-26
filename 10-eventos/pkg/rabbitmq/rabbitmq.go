package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel() (*amqp.Channel, error) {
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    if err != nil {
        panic(err)
    }
    rabbitMqChannel, err := conn.Channel()
    if err != nil {
        panic(err)
    }
    return rabbitMqChannel, nil
}

func Consume(rabbitMqChannel *amqp.Channel, out chan amqp.Delivery, queueName string) error{
	msgs, err := rabbitMqChannel.Consume(
		queueName,
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	
	if err != nil {
		return err
	}
	
	for msg := range msgs{
		out <- msg
	}

	return nil
}

func Publisher(rabbitMqChannel *amqp.Channel, body, exchangeName string) error{
	//  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    // defer cancel()
	println(body)
	err := rabbitMqChannel.PublishWithContext(
		context.Background(), 
		exchangeName,
		"",
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