package main

import (
	"log"
	"net/http"
	"net/url"
	"tat_gogogo/consts"
	"github.com/gin-gonic/gin"
)

func login(studentID string, password string) {

	res, err := http.PostForm(consts.Login,
		url.Values{"forceMobile": {"mobile"}, "mpassword": {password}, "muid": {studentID}})

	if err != nil {
		log.Fatalln(err)
	}

	res.Header.Set("Referer", consts.IndexPage)
	res.Header.Set("User-Agent", "Direk Android App")

	log.Println(res.Cookies())
}

func main() {

	log.Println(consts.Base)
	apiClient := gin.Default()

	apiClient.POST("/login", func(c *gin.Context) {
		studentId := c.PostForm("studentId")
		password := c.PostForm("password")

		login(studentId, password)

		c.JSON(200, gin.H{
			"status": "posted",
			"muid":   studentId,
		})
	})

	apiClient.Run() // listen and serve on 0.0.0.0:8080

}
