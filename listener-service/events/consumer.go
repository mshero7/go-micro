package events

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn      *amqp.Connection
	queueName string
}

func NewConsumer(conn *amqp.Connection) (Consumer, error) {
	consumer := Consumer{
		conn: conn,
	}

	err := consumer.setup()
	if err != nil {
		return Consumer{}, nil
	}

	return consumer, nil
}

func (c *Consumer) setup() error {
	// opens a unique, concurrent server channel to process the bulk of AMQP messages
	channel, err := c.conn.Channel()

	if err != nil {
		return err
	}

	return declareExchange(channel)
}

type Payload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (c *Consumer) Listen(topics []string) error {
	ch, err := c.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// we need to get random queue
	queue, err := declareRandomQueue(ch)
	if err != nil {
		return err
	}

	for _, s := range topics {
		ch.QueueBind(
			queue.Name,
			s,
			"logs_topic",
			false, // no-waits
			nil,
		)

		if err != nil {
			return err
		}
		fmt.Printf("Create Queue %s\n", queue.Name)
	}

	messages, err := ch.Consume(
		queue.Name,
		"",    //
		true,  // auto-acknowledge
		false, // exclusive
		false, // local
		false, // nowait
		nil,
	)
	if err != nil {

	}

	// consume all from RabbitMQ
	forever := make(chan bool)
	go func() {
		for d := range messages {
			var payload Payload
			_ = json.Unmarshal(d.Body, &payload)

			go handlePayload(payload)
		}
	}()

	fmt.Printf("Waiting for message [Exchange, Queue] [logs_topic, %s]\n", queue.Name)
	<-forever

	return nil
}

func handlePayload(payload Payload) {
	// We'll switch on the name value from the payload variable we received as a call to this function.
	switch payload.Name {
	case "log", "event":
		// log whatever we get
		err := logEvent(payload)
		if err != nil {
			log.Println(err)
		}

	case "auth":
		// authenticate

		// you can have as many cases as you want, as long as you write the logic
		// case "anything"
	default:
		err := logEvent(payload)
		if err != nil {
			log.Println(err)
		}
	}
}

func logEvent(entry Payload) error {
	jsonData, _ := json.MarshalIndent(entry, "", "\t")

	logServiceURL := "http://logger-service/log"

	req, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		return err
	}

	return nil
}
