package handler

import (
	"tat_gogogo/api/middleware"
	"tat_gogogo/glob/logs"
	"tat_gogogo/interface/controller"

	"github.com/gin-gonic/gin"
)

/*
LoginHandler is a function for gin to handle login api
*/
func LoginHandler(c *gin.Context) {
	authMiddleware, err := middleware.NewAuthMiddleware()
	if err != nil {
		logs.Warning.Printf("JWT Error:" + err.Error())
		c.Status(500)
	}

	studentID := c.PostForm("studentID")
	password := c.PostForm("password")

	loginController := controller.NewLoginController(studentID, password)
	result, err := loginController.Login()
	if err != nil {
		logs.Warning.Printf("LogIn failed:", studentID)
		c.Status(500)
		return
	}

	if result.GetStatus() != 200 {
		message := result.GetData()
		logs.Info.Println(message, "FROM", studentID)
		c.JSON(result.GetStatus(), gin.H{
			"message": message,
		})
		return
	}

	authMiddleware.LoginHandler(c)
}
