package messagebroker

import (
	"SwordHealth/pkg/config"
	"SwordHealth/pkg/messagebroker/dtos"
	"encoding/json"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

type MQProducer struct {
	queue   string
	connURL string
}

func NewMQProducer() *MQProducer {
	env := config.InitEnvironmentConfig()
	return &MQProducer{
		queue:   env.RMQ_QUEUE,
		connURL: env.RMQ_URL,
	}
}

func (mqp *MQProducer) Produce(techName, taskTitle string, techID, taskID int, performedAt time.Time) error {
	conn, err := amqp091.Dial(mqp.connURL)
	defer conn.Close()

	if err != nil {
		log.Printf("Could not connect to Rabbitmq! %s", err)
	}

	ch, err := conn.Channel()
	defer ch.Close()
	if err != nil {
		log.Printf("Could not open a channel! %s", err)
	}

	queue, err := ch.QueueDeclare(mqp.queue, false, false, false, false, nil)

	if err != nil {
		log.Printf("Could not declare queue! %s", err)
	}

	message := dtos.Message{
		TechName:    techName,
		TechID:      techID,
		TaskTitle:   taskTitle,
		TaskID:      taskID,
		PerformedAt: performedAt,
	}

	jsonMessage, err := json.Marshal(&message)
	if err != nil {
		log.Printf("Could not marshal message! %s\n", err)
		return err
	}

	if err := ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/json",
			Body:        []byte(jsonMessage),
		},
	); err != nil {
		log.Printf("Could not publish message to MQ! %s\n", err)
		return err
	}

	return nil
}
