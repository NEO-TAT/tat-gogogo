package service

import (
	"net/http"
	"tat_gogogo/domain/repository"
	"tat_gogogo/glob/logs"
	"tat_gogogo/utilities/httcli"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/viper"
)

/*
InfoService contains service of Result
*/
type InfoService struct {
	repo repository.InfoRepository
}

/*
NewInfoService init a new InfoService
*/
func NewInfoService(repo repository.InfoRepository) *InfoService {
	return &InfoService{repo: repo}
}

/*
GetInfoRows get ifno rows
*/
func (i *InfoService) GetInfoRows(studentID, year, semester string) (*goquery.Selection, error) {
	buffer, err := getCourseSelectResponse(studentID, year, semester)
	if err != nil {
		return nil, err
	}

	defer buffer.Body.Close()

	courseDoc, err := goquery.NewDocumentFromReader(buffer.Body)
	if err != nil {
		logs.Error.Panicln(err)
		return nil, err
	}

	return courseDoc.Find("table").Last().Find("tr"), nil
}

func getCourseSelectResponse(studentID, year, semester string) (buffer *http.Response, err error) {
	bufferReq, err := http.NewRequest(
		"GET",
		viper.GetString("COURESESYSTEM.Select"),
		nil)
	if err != nil {
		logs.Error.Panicln(err)
		return nil, err
	}

	bufferReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	q := bufferReq.URL.Query()
	q.Add("format", "-2")
	q.Add("code", studentID)
	q.Add("year", year)
	q.Add("sem", semester)

	bufferReq.URL.RawQuery = q.Encode()

	client := httcli.GetInstance()

	buffer, err = client.Do(bufferReq)
	if err != nil {
		logs.Error.Panicln(err)
		return nil, err
	}

	return buffer, nil
}
