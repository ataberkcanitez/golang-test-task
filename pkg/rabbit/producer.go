package rabbit

import (
	"bytes"
	"encoding/json"

	"github.com/streadway/amqp"
)

type Rabbit struct {
	Conn *amqp.Connection
}

type Message struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Message  string `json:"message"`
}

func NewRabbit() (*Rabbit, error) {
	conn, err := amqp.Dial("amqp://user:password@localhost:7000")
	if err != nil {
		return nil, err
	}

	return &Rabbit{Conn: conn}, nil
}

func (rabbit *Rabbit) Publish(channel *amqp.Channel, message Message) error {
	q, err := channel.QueueDeclare("catbyte", false, false, false, false, nil)
	if err != nil {
		return err
	}

	encodedBody, err := message.Encode()
	if err != nil {
		return err
	}

	err = channel.Publish("", q.Name, false, false, amqp.Publishing{ContentType: "text/plain", Body: encodedBody})
	if err != nil {
		return nil
	}

	return nil
}

func (message *Message) Encode() ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(message)
	return b.Bytes(), err

}
