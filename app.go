package main

import (
	"tat_gogogo/server"
	"github.com/gin-gonic/gin"
)

var HttpServer *gin.Engine

func main() {
	server.Run(HttpServer)
}
