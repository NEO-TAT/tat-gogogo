package portal

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"tat_gogogo/configs"
	"tat_gogogo/model"
)

/*
Login is a function which handle login request
return: http.Client for future reuse
studentID: the id of student
password: the password of student
*/
func Login(client *http.Client, studentID string, password string) (loginResult model.Result, err error) {
	req := newRequest(studentID, password)
	resp, err := client.Do(req)
	loginResult = handleResponse(resp)

	return loginResult, err
}

/*
NewClient is a function which init a http client for crawler
*/
func NewClient() *http.Client {
	cookieJar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: cookieJar,
	}

	return client
}

func handleResponse(resp *http.Response) (loginResult model.Result) {
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

	return *model.NewResult(isSuccess, statusCode, message)
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

	req, err := http.NewRequest("POST", config.Portal.Login, bytes.NewBufferString(data.Encode()))

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", config.Portal.IndexPage)
	req.Header.Set("User-Agent", "Direk Android App")

	return req
}
