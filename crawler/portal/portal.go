package portal

import (
	"bytes"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"tat_gogogo/configs"
)

/*
Login is a function which handle login request
return: http.Client for future reuse
studentID: the id of student
password: the password of student
*/
func Login(client *http.Client, studentID string, password string) (*http.Response, error) {
	req := newRequest(studentID, password)
	resp, err := client.Do(req)

	return resp, err
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
