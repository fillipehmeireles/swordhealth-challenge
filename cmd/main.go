package main

import (
	"SwordHealth/cmd/api"
	"SwordHealth/cmd/messagebroker"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go api.ApiListener(&waitGroup)
	go messagebroker.MessageBrokerListener()
	waitGroup.Wait()
}
