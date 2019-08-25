package login

import (
	"log"

	"tat_gogogo/infrastructure/api/handler"
	"tat_gogogo/infrastructure/middleware"

	"github.com/gin-gonic/gin"
)

/*
Controller is a function for gin to handle login api
*/
func Controller(c *gin.Context) {
	authMiddleware, err := middleware.NewAuthMiddleware()
	if err != nil {
		c.Status(500)
		log.Fatal("JWT Error:" + err.Error())

	}

	studentID := c.PostForm("studentID")
	password := c.PostForm("password")

	loginHandler := handler.NewLoginHanlder(studentID, password)
	result, err := loginHandler.Login()

	if err != nil {
		log.Panicln("failed to fetch login cookie")
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
