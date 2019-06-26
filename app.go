package main

import (
	"log"
	"tat_gogogo/consts"
	"tat_gogogo/routers"
)


func main() {

	log.Println(consts.Base)
	apiClient := routers.Init()

	apiClient.Run() // listen and serve on 0.0.0.0:8080

}
