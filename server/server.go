package server

import(
	"github.com/gin-gonic/gin"
	"tat_gogogo/routes"
)

func Run(httpServer *gin.Engine) {

	httpServer = gin.Default()

	routes.RegisterRoutes(httpServer)

	// listen and serve on 0.0.0.0:8080
	err := httpServer.Run()

	if err != nil {
		panic("server run error: " + err.Error())
	}
}