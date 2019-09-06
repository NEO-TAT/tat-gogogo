package service

import (
	"os"
	"os/signal"
	"syscall"
	"tat_gogogo/utilities/logs"

	"github.com/spf13/viper"
)

/*
Start is init system after server start then wait stop signal.
*/
func Start() {
	configInit()

	go logs.LogInit()

	go serviceStart()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill, syscall.SIGTERM)
	for {
		select {
		case <-stop:
			close()
			panic(nil)
		}
	}
}

/*
serviceStart is create gin engine and running.
*/

func serviceStart() {
	router := ginInit()
	logs.Error.Panicln(router.Run(":" + viper.GetString("PORT")))
}
