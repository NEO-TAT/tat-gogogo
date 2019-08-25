package main

import (
	"log"
	"tat_gogogo/configs"
	"tat_gogogo/infrastructure/router"

	"github.com/gin-gonic/gin"
)

func run() {
	var httpServer *gin.Engine

	configuration, err := configs.New()

	if err != nil {
		log.Panicln("Configuration err", err)
	}

	httpServer = gin.Default()

	router.Register(httpServer)

	serverAddr := configuration.Constants.Host + ":" + configuration.Constants.Port

	// listen and serve on 0.0.0.0:8080
	httpServer.Run(serverAddr)
}

func main() {
	run()
}
