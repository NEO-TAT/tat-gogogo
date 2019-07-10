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
HandleRequest is a function which handles the request
*/
func HandleRequest(c *gin.Context) {

	studentID := c.PostForm("studentId")
	password := c.PostForm("password")
	_, resp := Login(studentID, password)

	handleResponse(c, resp)
}

/*
Login is a function which handle login request
return: http.Client for future reuse
studentID: the id of student
password: the password of student
*/
func Login(studentID string, password string) (*http.Client, *http.Response) {
	client := newClient()
	req := newRequest(studentID, password)

	resp, _ := client.Do(req)

	return client, resp
}

func handleResponse(c *gin.Context, resp *http.Response) {
	defer resp.Body.Close()

	var data map[string]interface{}
	var result Result
	json.NewDecoder(resp.Body).Decode(&data)

	statusCode := 200
	isSuccess := data["success"].(bool)
	message := "登入成功"
	if !isSuccess {
		statusCode = 401
		message = "帳號或密碼錯誤，請重新輸入。"
	}

	result = Result{success: isSuccess, status: statusCode, message: message}

	if result.status == 200 {
		c.Status(200)
	} else {
		c.JSON(result.status, gin.H{
			"message": result.message,
		})
	}
}

func newClient() *http.Client {
	cookieJar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: cookieJar,
	}

	return client
}

func newRequest(studentID string, password string) *http.Request {
	config, err := configs.New()
	if err != nil {
		log.Panicln("failed to new configuration")
	}

	data := url.Values{
		"forceMobile": {"mobile"},
		"mpassword":   {password},
		"muid":        {studentID},
	}

	req, err := http.NewRequest("POST", config.PORTAL.Login, bytes.NewBufferString(data.Encode()))

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", config.PORTAL.IndexPage)
	req.Header.Set("User-Agent", "Direk Android App")

	return req
}
