package rabbitmq

import (
	"context"
	"sync"
	"time"

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

	err = rabbitMqChannel.QueueBind(
		"green-q",        // queue name
		"green",             // routing key
		"amq.direct", // exchange
		false,
		nil,
	)

		err = rabbitMqChannel.QueueBind(
		"green-q",        // queue name
		"black",             // routing key
		"amq.direct", // exchange
		false,
		nil,
	)

	if err != nil{
		panic(err)
	}

		err = rabbitMqChannel.QueueBind(
		"blue-q",        // queue name
		"blue",             // routing key
		"amq.direct", // exchange
		false,
		nil,
	)

	if err != nil{
		panic(err.Error())
	}


    return rabbitMqChannel, nil
}

func Consume(rabbitMqChannel *amqp.Channel, out chan amqp.Delivery, queueName string, wg *sync.WaitGroup) error {
		msgs, err := rabbitMqChannel.Consume(
			queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
		)

		if err != nil {
			println(string(err.Error()))
			return err 
		}
		for msg := range msgs{
			println("a")
			out <- msg
		}

		defer wg.Done()
	return nil
}

func Publisher(rabbitMqChannel *amqp.Channel, body, exchangeName, keyName string, wg *sync.WaitGroup) error{
	 ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
	err := rabbitMqChannel.PublishWithContext(
		ctx, 
		exchangeName,
		keyName,
		false,
		false,
		amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(body),
        },
	)
	defer wg.Done()
	if err != nil {
		return err
	}

	return nil
}