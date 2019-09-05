package service

import (
	"os"
	"os/signal"
	"syscall"
	"tat_gogogo/glob/logs"

	"github.com/spf13/viper"
)

func Start() {
	// Init
	configInit()
	logs.LogInit()
	// Server Start
	go serviceStart()
	// Server Safe Stop
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill, syscall.SIGTERM)
	for {
		select {
		case <-stop:
			panic(nil)
		}
	}
}

func serviceStart() {
	router := ginInit()
	logs.Error.Panicln(router.Run(":" + viper.GetString("PORT")))
}

// router.RunTLS(
// 	":"+viper.GetInt("PORT"),
// 	"./SSL/server.crt",
// "./SSL/server.key",)