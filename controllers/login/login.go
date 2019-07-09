package login

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"tat_gogogo/configs"

	"github.com/gin-gonic/gin"
)

/*
Result is the result of Login response
success: is login successed
status: the status of response
message: show user that how login is going on
*/
type Result struct {
	success bool
	status  int
	message string
}

/*
Controller is a struct which manipulate login request and send result back to user
studentID: the id of student
password: the password of student
*/
type Controller struct {
	studentID string
	password  string
}

/*
HandleLogin is a function for gin to handle login api
*/
func HandleLogin(c *gin.Context) {
	studentID := c.PostForm("studentId")
	password := c.PostForm("password")

	controller := Controller{studentID: studentID, password: password}

	result := controller.handleRequest()

	c.JSON(result.status, gin.H{
		"success": result.success,
		"message": result.message,
	})
}

func (controller *Controller) newClient() (*http.Client, *http.Request) {
	config, err := configs.New()
	if err != nil {
		log.Panicln("failed to new configuration")
	}

	cookieJar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: cookieJar,
	}

	data := url.Values{
		"forceMobile": {"mobile"},
		"mpassword":   {controller.password},
		"muid":        {controller.studentID},
	}

	req, err := http.NewRequest("POST", config.PORTAL.Login, bytes.NewBufferString(data.Encode()))

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", config.PORTAL.IndexPage)
	req.Header.Set("User-Agent", "Direk Android App")

	return client, req
}

func (controller *Controller) handleRequest() Result {
	client, req := controller.newClient()
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
		return Result{success: false, status: 401}
	}

	defer resp.Body.Close()

	var data map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&data)

	statusCode := 200
	isSuccess := data["success"].(bool)
	message := "登入成功"
	if !isSuccess {
		statusCode = 401
		message = "帳號或密碼錯誤，請重新輸入。"
	}

	return Result{success: isSuccess, status: statusCode, message: message}
}
