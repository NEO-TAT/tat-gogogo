package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/cookieJar"
	"net/url"
	"tat_gogogo/consts"
	"log"
	"bytes"
	//"io/ioutil"
)

type Result struct {
	success bool
	status int
}

func LoginController(c *gin.Context) {
	studentId := c.PostForm("studentId")
	password := c.PostForm("password")

	login(studentId, password)

	c.JSON(200, gin.H{
		"status": "posted",
		"muid":   studentId,
	})
}

func newClient(studentID string, password string) (*http.Client, *http.Request) {
	cookieJar, _ := cookiejar.New(nil)
	
	client := &http.Client{
		Jar: cookieJar,
	}

	data := 	url.Values{
		"forceMobile": {"mobile"},
		"mpassword": {password}, 
		"muid": {studentID},
	}

	req, err := http.NewRequest("POST", consts.Login, bytes.NewBufferString(data.Encode()))

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", consts.IndexPage)
	req.Header.Set("User-Agent", "Direk Android App")

	return client, req
}

func handleRequest(studentID string, password string) (Result) {
	client, req := newClient(studentID, password)
	resp, err := client.Do(req)
	
	if err != nil {
		log.Fatalln(err)
		return Result{success: false, status: 401}
	}

	defer resp.Body.Close()

	return Result{success: true, status: 200}
}

func login(studentID string, password string) {

	result := handleRequest(studentID, password)

	log.Println(result)
}