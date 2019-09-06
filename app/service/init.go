package service

import (
	"os"
	"tat_gogogo/infrastructure/router"
	"tat_gogogo/utilities/logs"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

/*
configInit is init load config.
*/
func configInit() {
	viper.SetConfigName("config")
	viper.AddConfigPath("..")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Fatal error config file: " + err.Error())
	}
}

/*
ginInit is create gin engine and register middleware.
*/
func ginInit() *gin.Engine {
	ginRouter := gin.Default()

	CORS := cors.DefaultConfig()
	CORS.AllowAllOrigins = true
	CORS.AllowCredentials = true
	CORS.AllowWebSockets = true
	ginRouter.Use(cors.New(CORS))

	pprof.Register(ginRouter)

	logFile, err := os.Create("./log/restful_server.log")
	if err != nil {
		logs.Warning.Println(err)
	} else {
		ginRouter.Use(gin.LoggerWithWriter(logFile))
	}
	router.Register(ginRouter)

	return ginRouter
}
