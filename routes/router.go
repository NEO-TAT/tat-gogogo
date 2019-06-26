package routes

import (
	"github.com/gin-gonic/gin"
	"tat_gogogo/controllers/login"
)

func RegisterRoutes(router *gin.Engine) {

	router.POST("/login", login.LoginController)

}