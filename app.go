package main

import (
	"log"
	"tat_gogogo/service"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
	service.Start()
}
