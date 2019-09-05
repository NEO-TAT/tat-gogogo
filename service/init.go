package service

import (
	"os"
	"tat_gogogo/glob/logs"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func configInit() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Fatal error config file: " + err.Error())
	}
}

func ginInit() *gin.Engine {
	router := gin.Default()
	CORS := cors.DefaultConfig()
	CORS.AllowAllOrigins = true
	CORS.AllowCredentials = true
	CORS.AllowWebSockets = true
	router.Use(cors.New(CORS))
	pprof.Register(router)
	// -----------------------------------------------[Log]
	logFile, err := os.Create("./log/restful_server.log")
	if err != nil {
		logs.Warning.Println(err)
	} else {
		router.Use(gin.LoggerWithWriter(logFile))
	}
	// -----------------------------------------------[return]
	return router
}
