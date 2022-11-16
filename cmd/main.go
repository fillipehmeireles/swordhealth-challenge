package main

import (
	"SwordHealth/cmd/api"
	"SwordHealth/cmd/messagebroker"
	"SwordHealth/pkg/api/drivers"
	"SwordHealth/pkg/config"
	"log"
	"sync"
)

func migrationManager(waitGroup *sync.WaitGroup) {

	defer waitGroup.Done()
	if migrationEnv := config.InitEnvironmentConfig().MIGRATE; migrationEnv {
		dbDriver := drivers.NewDBDriver()
		dbDriver.Connect()
		defer dbDriver.Close()
		log.Println("[i] Migration is ON")
		dbDriver.Migrate()
	}

}
func main() {

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go migrationManager(&waitGroup)
	//waitGroup.Wait()
	go api.ApiListener(&waitGroup)
	go messagebroker.MessageBrokerListener()
	waitGroup.Wait()

}
