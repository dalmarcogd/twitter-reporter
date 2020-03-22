package rabbit

import (
	"fmt"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-processor/brokers/events"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-processor/environments"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-processor/errors"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-processor/utils"
	"github.com/streadway/amqp"
)

func GetConnection() (*amqp.Connection, error) {
	env := environments.GetEnvironment()
	return amqp.Dial(fmt.Sprintf("amqp://%v:%v@%v:%v/%v", env.RabbitUsername, env.RabbitPassword, env.RabbitURL, env.RabbitPort, env.RabbitVHost))
}

func Publish(event events.Event) error {
	connection, err := GetConnection()
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
