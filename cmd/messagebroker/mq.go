package messagebroker

import "SwordHealth/pkg/messagebroker"

func MessageBrokerListener() {
	mqc := messagebroker.NewMQConsumer()
	mqc.Consume()
}
