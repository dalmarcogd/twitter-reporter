package rabbit

import (
	"context"
	"fmt"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/brokers/events"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/environments"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/errors"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/utils"
	"github.com/streadway/amqp"
)

func getConnection() (*amqp.Connection, error) {
	env := environments.GetEnvironment()
	return amqp.Dial(fmt.Sprintf("amqp://%v:%v@%v:%v/%v", env.RabbitUsername, env.RabbitPassword, env.RabbitURL, env.RabbitPort, env.RabbitVHost))
}

func Publish(ctx context.Context, event events.Event) error {
	connection, err := getConnection()
	if err != nil {
		return errors.NewRabbitConnectionError(err)
	}
	defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		return err
	}

	body, err := utils.NewJsonConverter().Encode(event)
	if err != nil {
		return err
	}

	fmt.Print(string(body))

	if err = channel.ExchangeDeclare(event.GetName(),
		amqp.ExchangeTopic,
		true,
		false,
		false,
		false,
		nil); err != nil {
		return err
	}

	return channel.Publish(event.GetName(), "*", false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Body:         body,
	})
}
