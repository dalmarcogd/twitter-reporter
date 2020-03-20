package rabbit

import (
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

type rabbit struct{}

func NewRabbit() *rabbit {
	return &rabbit{}
}

func (r rabbit) Publish(event events.Event) error {
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

	queue, err := channel.QueueDeclare(
		"accounts-persist",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return channel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Body:         body,
	})
}
