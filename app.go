package main

import (
	"github.com/gin-gonic/gin"
	"tat_gogogo/configs"
	"tat_gogogo/interface/routes"
	"log"
)

func run() {
	var httpServer *gin.Engine

	configuration, err := configs.New()

	if err != nil {
		log.Panicln("Configuration err", err)
	}

	httpServer = gin.Default()

	routes.RegisterRoutes(httpServer)

	serverAddr := configuration.Constants.Host + ":" + configuration.Constants.Port

	// listen and serve on 0.0.0.0:8080
	httpServer.Run(serverAddr)
}

func main() {
	run()
}
