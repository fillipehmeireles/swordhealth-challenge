package main

import (
	"SwordHealth/pkg/config"
	"SwordHealth/pkg/drivers"
	"SwordHealth/pkg/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func migrationManager() {
	if migrationEnv := config.InitEnvironmentConfig().MIGRATE; migrationEnv {
		dbDriver := drivers.NewDBDriver()
		dbDriver.Connect()
		defer dbDriver.Close()
		log.Println("[i] Migration is ON")
		dbDriver.Migrate()
	}
}
func main() {
	migrationManager()
	log.Printf("Hello World! %s", config.InitEnvironmentConfig().DB_URL)
	r := gin.Default()
	routes.Routes(r)
	log.Println("Server is running")
	r.Run()
}
