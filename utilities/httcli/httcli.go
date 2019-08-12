package httcli

import (
	"net/http"
	"net/http/cookiejar"
)

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
