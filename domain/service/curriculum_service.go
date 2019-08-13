package service

import (
	"log"
	"net/http"
	"net/url"
	"strings"
	"tat_gogogo/configs"
	"tat_gogogo/domain/model"
	"tat_gogogo/domain/repository"
	"tat_gogogo/utilities/decoder"
	"tat_gogogo/utilities/httcli"

	"github.com/PuerkitoBio/goquery"
)

var (
	config, configError = configs.New()
	client              = httcli.GetInstance()
	columnMap           = map[int]string{
		0:  "id",
		1:  "name",
		6:  "instructor",
		8:  "periodsOfSunday",
		9:  "periodsOfMonday",
		10: "periodsOfTuesday",
		11: "periodsOfWednesday",
		12: "periodsOfThursday",
		13: "periodsOfFriday",
		14: "periodsOfSaturday",
		15: "classroom",
	}
)

type curriculum model.Curriculum

/*
CurriculumService contains service of curriculum
*/
type CurriculumService struct {
	repo repository.CurriculumRepository
}

/*
NewCurriculumService init a new CurriculumService
@paramter: repository.CurriculumRepository)
@return: *CurriculumService
*/
func NewCurriculumService(repo repository.CurriculumRepository) *CurriculumService {
	return &CurriculumService{repo: repo}
}

/*
IsLoginCurriculum judje if curriculum login successful
@return: bool, error
*/
func (c *CurriculumService) IsLoginCurriculum() (bool, error) {
	doc, err := postSSOLoginCourseSystem()
	if err != nil {
		log.Panicln(err)
		return false, err
	}

	isAccessCourse, err := isAccessCourse(doc)
	if err != nil {
		log.Panicln(err)
		return false, err
	}

	return isAccessCourse, nil
}

/*
GetCurriculumDocument will get curriculum doc from the NewRequest
@paramter: targetStudentID string
@return: *goquery.Document, error
*/
func (c *CurriculumService) GetCurriculumDocument(targetStudentID string) (*goquery.Document, error) {
	form := url.Values{
		"code":   {targetStudentID},
		"format": {"-3"},
	}
	curriculumRequest, err := http.NewRequest("POST", config.CoureseSystem.Select, strings.NewReader(form.Encode()))
	if err != nil {
		log.Panicln(err)
		return nil, err
	}
	curriculumRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	curriculumsResp, err := client.Do(curriculumRequest)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}
	curriculumDoc, err := goquery.NewDocumentFromReader(curriculumsResp.Body)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	return curriculumDoc, nil
}

func postSSOLoginCourseSystem() (*goquery.Document, error) {
	req, err := http.NewRequest("POST", config.Portal.SsoLoginCourseSystem, nil)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	return doc, nil
}

func isAccessCourse(doc *goquery.Document) (bool, error) {
	form := parseFormBy(doc)

	bufferRequest, err := http.NewRequest("POST", config.CoureseSystem.MainPage, strings.NewReader(form.Encode()))
	if err != nil {
		log.Panicln(err)
		return false, err
	}

	bufferRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	buffer, err := client.Do(bufferRequest)
	if err != nil {
		log.Panicln(err)
		return false, err
	}

	return isLoginCurriculumSuccess(buffer)
}

func parseFormBy(doc *goquery.Document) url.Values {
	form := url.Values{}
	doc.Find("input[type=hidden]").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("name")
		value, _ := s.Attr("value")
		form.Add(name, value)
	})

	return form
}

func isLoginCurriculumSuccess(buffer *http.Response) (bool, error) {
	defer buffer.Body.Close()

	courseDoc, err := goquery.NewDocumentFromReader(buffer.Body)
	if err != nil {
		log.Panicln(err)
		return false, err
	}

	rawLast := courseDoc.Find("body a").Last().Text()
	last, err := decoder.DecodeToBig5(rawLast)
	if err != nil {
		log.Println(err)
		return false, err
	}

	return last == "依 [學號]／[課號] 查詢選課表", nil
}
