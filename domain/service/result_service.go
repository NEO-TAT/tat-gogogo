package service

import (
	"bytes"
	"net/http"
	"net/url"

	"tat_gogogo/domain/repository"
	"tat_gogogo/glob/logs"

	"github.com/spf13/viper"
)

/*
ResultService contains service of Result
*/
type ResultService struct {
	repo repository.ResultRepository
}

/*
NewResultService init a new ResultService
*/
func NewResultService(repo repository.ResultRepository) *ResultService {
	return &ResultService{repo: repo}
}

/*
NewLoginRequest init a login request
*/
func (r *ResultService) NewLoginRequest(studentID, password string) *http.Request {
	data := url.Values{
		"forceMobile": {"mobile"},
		"mpassword":   {password},
		"muid":        {studentID},
	}

	req, err := http.NewRequest(
		"POST",
		viper.GetString("PORTAL.Login"),
		bytes.NewBufferString(data.Encode()))

	if err != nil {
		logs.Error.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", viper.GetString("PORTAL.IndexPage"))
	req.Header.Set("User-Agent", "Direk Android App")

	return req
}
