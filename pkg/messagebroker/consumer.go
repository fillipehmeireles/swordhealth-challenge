package messagebroker

import (
	"SwordHealth/pkg/config"
	"SwordHealth/pkg/messagebroker/dtos"
	"encoding/json"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type MQConsumer struct {
	queue   string
	connURL string
}

func NewMQConsumer() *MQConsumer {
	env := config.InitEnvironmentConfig()
	log.Println(env.RMQ_URL)
	return &MQConsumer{
		queue:   env.RMQ_QUEUE,
		connURL: env.RMQ_URL,
	}
}

func (mqc *MQConsumer) Consume() {
	conn, err := amqp091.Dial(mqc.connURL)
	defer conn.Close()

	if err != nil {
		log.Printf("Could not connect to Rabbitmq! %s", err)
	}

	ch, err := conn.Channel()
	defer ch.Close()
	if err != nil {
		log.Printf("Could not open a channel! %s", err)
	}

	queue, err := ch.QueueDeclare(mqc.queue, false, false, false, false, nil)

	if err != nil {
		log.Printf("Could not declare queue! %s", err)
	}

	msgs, err := ch.Consume(queue.Name, "", false, false, false, false, nil)

	if err != nil {
		log.Printf("Could not create a consumer! %s", err)
	}

	forever := make(chan bool)

	go func() {
		log.Println("RMQ Consumer is listening.")
		for message := range msgs {
			var msg dtos.Message
			if err := json.Unmarshal(message.Body, &msg); err != nil {
				log.Printf("Could not unmarshal MQ message body! %s\n", err)
			} else {
				log.Printf("[MQ MESSAGE] The tech %s (%d) performed the task %s (%d) on date (%s)", msg.TechName, msg.TechID, msg.TaskTitle, msg.TaskID, msg.PerformedAt)
			}
		}
	}()

	<-forever

}
