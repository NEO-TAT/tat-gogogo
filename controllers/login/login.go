package login

import (
	"tat_gogogo/crawler/portal"

	"log"

	"github.com/gin-gonic/gin"
)

/*
Controller is a function for gin to handle login api
*/
func Controller(c *gin.Context) {
	studentID := c.PostForm("studentID")
	password := c.PostForm("password")
	client := portal.NewClient()

	result, err := portal.Login(client, studentID, password)
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

	c.Status(200)

}
