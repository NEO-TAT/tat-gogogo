package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"tat_gogogo/consts"
	"log"
)

func LoginController(c *gin.Context) {
	studentId := c.PostForm("studentId")
	password := c.PostForm("password")

	login(studentId, password)

	c.JSON(200, gin.H{
		"status": "posted",
		"muid":   studentId,
	})
}

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