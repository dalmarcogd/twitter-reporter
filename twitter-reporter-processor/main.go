package main

import (
	"context"
	"fmt"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-processor/brokers/events"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-processor/brokers/rabbit"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-processor/database"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-processor/errors"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-processor/monitoring"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-processor/services/twitter"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-processor/utils"
	"go.elastic.co/apm"
	"log"
	"strconv"
	"time"
)

func main() {
	errors.FailOnError(database.Migrate(), "Fail on migrate database")
	defer database.CloseConnection()

	connection, err := rabbit.GetConnection()
	errors.FailOnError(err, "Error when get connection")
	defer func() {
		errors.FailOnError(connection.Close(), "Error on close connection with rabbit")
	}()

	channel, err := connection.Channel()
	errors.FailOnError(err, "Error when get channel")
	defer func() {
		errors.FailOnError(channel.Close(), "Error when close channel")
	}()

	err = channel.Qos(5, 0, false)
	errors.FailOnError(err, "Error setup qos")

	queue, err := channel.QueueDeclare(
		"twitter-reporter-processor", // name
		true,                         // durable
		false,                        // delete when unused
		false,                        // exclusive
		false,                        // no-wait
		nil,                          // arguments
	)
	errors.FailOnError(err, "Error when declare a queue")

	errors.FailOnError(channel.QueueBind(queue.Name, "*", "ReporterEvent", false, nil), "Error when bind with exchange")

	forever := make(chan bool)

	// Workers
	for i := 0; i < 5; i++ {
		go func(i int) {
			log.Printf(fmt.Sprintf("Start consumer worker: %s", strconv.Itoa(i)))

			msgs, err := channel.Consume(queue.Name,
				"",
				false,
				false,
				false,
				false,
				nil)
			errors.FailOnError(err, "Error when create consumer")

			for msg := range msgs {
				reporterEvent := events.NewReporterEvent("", "")

				tx := monitoring.GetTracer().StartTransaction(fmt.Sprintf("Consuming message %s", reporterEvent.GetName()), "amqp.consumer")
				ctx := apm.ContextWithTransaction(context.Background(), tx)

				errors.PrintOnError(utils.NewJsonConverter().Decode(msg.Body, reporterEvent), "Error when decode message")

				log.Printf("Start process reporter %s:%s", reporterEvent.ReporterId, reporterEvent.Tag)

				errors.PrintOnError(database.CreateReporter(ctx, reporterEvent.ReporterId, reporterEvent.Tag), "Error when save account")
				log.Printf("Reporter saved %s", reporterEvent.ReporterId)

				tweets, err := twitter.GetTweetsByHashtag(ctx, reporterEvent.Tag)
				errors.PrintOnError(err, "Error when get messages from twitter")

				log.Printf("Start persist tweets(%s)", strconv.Itoa(len(tweets)))
				_ = utils.SpanTracer(ctx, fmt.Sprintf("Start persist tweets(%s)", strconv.Itoa(len(tweets))), "function", func(cx context.Context, span *apm.Span) error {
					for _, t := range tweets {
						tweet := t.(map[string]interface{})

						var twitterUser database.TwitterUserModel

						if tUser, ok := tweet["user"]; ok {
							tweetUser := tUser.(map[string]interface{})
							tweetUserId := tweetUser["id"].(float64)
							tweetUserName := tweetUser["name"].(string)
							tweetUserScreenName := tweetUser["screen_name"].(string)
							tweetUserStatusesCount := tweetUser["statuses_count"].(float64)
							tweetUserFollowersCount := tweetUser["followers_count"].(float64)
							tweetUserLocation := tweetUser["location"].(string)

							twitterUser, err = database.GetTwitterUserById(ctx, tweetUserId)
							if err != nil {
								twitterUser, err = database.CreateTwitterUser(ctx, tweetUserId, tweetUserName, tweetUserScreenName, tweetUserStatusesCount, tweetUserFollowersCount, tweetUserLocation)
								errors.PrintOnError(err, "Error when save user")
								if err != nil {
									continue
								}
							} else {
								twitterUser, err = database.UpdateTwitterUser(ctx, tweetUserId, tweetUserName, tweetUserScreenName, tweetUserStatusesCount, tweetUserFollowersCount, tweetUserLocation)
								errors.PrintOnError(err, "Error when save user")
								if err != nil {
									continue
								}
							}

							tweetId := tweet["id"].(float64)
							tweetText := tweet["text"].(string)
							tweetLanguage := ""
							if v, ok := tweet["lang"]; ok {
								tweetLanguage = v.(string)
							}
							tweetCreatedAtStr := tweet["created_at"].(string)
							tweetCreatedAt, err := time.Parse(time.RubyDate, tweetCreatedAtStr)
							errors.PrintOnError(err, "Error when parse created_at")
							if err != nil {
								continue
							}

							_, err = database.GetTwitterTweetById(ctx, tweetUserId)
							if err != nil {
								err := database.CreateTwitterTweet(ctx, tweetId, reporterEvent.ReporterId, twitterUser.Id, tweetText, tweetLanguage, tweetCreatedAt)
								errors.PrintOnError(err, "Error when save tweet")
								if err != nil {
									continue
								}
							}
						} else {
							log.Print("The Tweet cannot be processed because the user was not found")
						}
					}
					return err
				})

				log.Print("End persist tweets")

				tx.End()
				ctx.Done()
				log.Printf("End process reporter %s:%s", reporterEvent.ReporterId, reporterEvent.Tag)
				errors.PrintOnError(msg.Ack(false), "Error when ack the message")
			}
		}(i)
	}
	<-forever
}
