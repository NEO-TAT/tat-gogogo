package handler

import (
	"log"

	"tat_gogogo/infrastructure/middleware"
	"tat_gogogo/interface/controller"

	"github.com/gin-gonic/gin"
)

/*
LoginHandler is a function for gin to handle login api
*/
func LoginHandler(c *gin.Context) {
	authMiddleware, err := middleware.NewAuthMiddleware()
	if err != nil {
		log.Printf("JWT Error:" + err.Error())
		c.Status(500)
	}

	studentID := c.PostForm("studentID")
	password := c.PostForm("password")

	loginController := controller.NewLoginController(studentID, password)
	result, err := loginController.Login()

	if err != nil {
		log.Printf("failed to fetch login cookie")
		c.Status(500)
		return
	}

	if result.GetStatus() != 200 {
		c.JSON(result.GetStatus(), gin.H{
			"message": result.GetData(),
		})
		return
	}

	authMiddleware.LoginHandler(c)
}
