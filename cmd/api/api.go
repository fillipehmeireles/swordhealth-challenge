package api

import (
	"SwordHealth/pkg/api/routes"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
)

func ApiListener(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	r := gin.Default()
	routes.Routes(r)
	log.Println("Server is running")
	r.Run()
}
