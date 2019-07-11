package login

import (
	"encoding/json"
	"net/http"
	"tat_gogogo/crawler/portal"

	"log"

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
Controller is a function for gin to handle login api
*/
func Controller(c *gin.Context) {
	studentID := c.PostForm("studentId")
	password := c.PostForm("password")
	client := portal.NewClient()
	resp, err := portal.Login(client, studentID, password)

	if err != nil {
		log.Panicln("failed to fetch login cookie")
		c.Status(500)
		return
	}

	handleResponse(c, resp)
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
