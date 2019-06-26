package server

import(
	"github.com/gin-gonic/gin"
	"tat_gogogo/routes"
	"tat_gogogo/configs"
)

func Run(httpServer *gin.Engine) {

	serverConfig := configs.GetServerConfig()

	httpServer = gin.Default()

	routes.RegisterRoutes(httpServer)

	serverAddr := serverConfig["HOST"] + ":" + serverConfig["PORT"]

	// listen and serve on 0.0.0.0:8080
	err := httpServer.Run(serverAddr)

	if err != nil {
		panic("server run error: " + err.Error())
	}
}