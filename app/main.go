package main

import (
	"log"
	"runtime/debug"
	"tat_gogogo/app/service"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(debug.Stack())
		}
	}()
	service.Start()
}
