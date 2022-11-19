package api

import (
	"SwordHealth/pkg/api/drivers"
	"SwordHealth/pkg/api/routes"
	"SwordHealth/pkg/config"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
)

func migrationManager() {
	if migrationEnv := config.InitEnvironmentConfig().MIGRATE; migrationEnv {
		dbDriver := drivers.NewDBDriver()
		dbDriver.Connect()
		defer dbDriver.Close()
		log.Println("[i] - Migration is ON")
		dbDriver.Migrate()
	}
}
func ApiListener(waitGroup *sync.WaitGroup) {
	migrationManager()
	defer waitGroup.Done()
	r := gin.Default()
	routes.Routes(r)
	log.Println("[i] - Server is running")
	r.Run()
}
