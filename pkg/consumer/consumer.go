package consumer

import (
	"encoding/json"
	"twitch_chat_analysis/pkg/redis"

	"github.com/streadway/amqp"
)

type Message struct {
	Sender   string `json:"sender"`
	Message  string `json:"message"`
	Receiver string `json:"receiver"`
}

type Rabbit struct {
	Conn *amqp.Connection
}

func NewRabbit(address string) (*Rabbit, error) {
	conn, err := amqp.Dial(address)
	if err != nil {
		return nil, err
	}

	return &Rabbit{
		Conn: conn,
	}, nil
}

func Consume() {
	rabbit, err := NewRabbit("amqp://user:password@localhost:7000")
	if err != nil {
		panic(err)
	}
	defer rabbit.Conn.Close()

	rdsClient := redis.NewRedis("localhost:6379")

	channel, err := rabbit.Conn.Channel()
	if err != nil {
		panic(err)
	}

	msgs, err := channel.Consume("catbyte", "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	var forever chan struct{}
	go func() {
		for msg := range msgs {
			var message Message
			json.Unmarshal(msg.Body, &message)

			err := rdsClient.LPush(message.Sender+":::"+message.Receiver, message.Message).Err()
			if err != nil {
				panic(err)
			}
		}
	}()

	<-forever

}
