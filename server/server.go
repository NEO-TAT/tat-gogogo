package server

import (
	"log"
	"tat_gogogo/configs"
	"tat_gogogo/routes"

	"github.com/gin-gonic/gin"
)

/*
Run is the enter point of project
*/
func Run(httpServer *gin.Engine) {

	configuration, err := configs.New()

	if err != nil {
		log.Panicln("Configuration err", err)
	}

	httpServer = gin.Default()

	routes.RegisterRoutes(httpServer)

	serverAddr := configuration.Constants.HOST + ":" + configuration.Constants.PORT

	// listen and serve on 0.0.0.0:8080
	httpServer.Run(serverAddr)
}
