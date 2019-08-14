package httcli

import (
	"net/http"
	"net/http/cookiejar"
	"sync"
)

var instance *http.Client
var once sync.Once

/*
GetInstance will get singleton instance
*/
func GetInstance() *http.Client {
	client := newClient()
	once.Do(func() {
		instance = client
	})
	return instance
}

func newClient() *http.Client {
	cookieJar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: cookieJar,
	}

	return client
}
