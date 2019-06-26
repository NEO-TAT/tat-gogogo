package routers

import (
	"github.com/gin-gonic/gin"
	"tat_gogogo/controllers/login"
)

func Init() *gin.Engine{
	apiClient := gin.Default()

	apiClient.POST("/login", login.LoginController)

	return apiClient
}