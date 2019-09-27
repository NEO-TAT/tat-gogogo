package service

import (
	"bytes"
	"log"
	"net/http"

	"net/url"
	"tat_gogogo/configs"
	"tat_gogogo/domain/repository"
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
