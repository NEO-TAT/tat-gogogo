package usecase

import (
	"bytes"
	"log"
	"net/http"
	"net/url"
	"tat_gogogo/configs"
	"tat_gogogo/domain/model"
	"tat_gogogo/domain/repository"
	"tat_gogogo/domain/service"
)

/*
ResultUsecase contains the functions for result usecase
*/
type ResultUsecase interface {
	GetService() *service.ResultService
	GetRepo() repository.ResultRepository
	LoginResult(client *http.Client, studentID, password string) (loginResult model.Result, err error)
	CurriculumResultBy(curriculumUsecase CurriculumUsecase, studentID, targetStudentID, year, semester string) (curriculumResult model.Result, err error)
	InfoResultBy(infoUsecase InfoUsecase, studentID, targetStudentID, year, semester string) (curriculumResult model.Result, err error)
}

type resultUsecase struct {
	repo    repository.ResultRepository
	service *service.ResultService
}

/*
NewResultUsecase init a new result usecase
@parameter: repository.ResultRepository, *service.ResultService
@return: *resultUsecase
*/
func NewResultUsecase(repo repository.ResultRepository, service *service.ResultService) *resultUsecase {
	return &resultUsecase{repo: repo, service: service}
}

/*
GetService get usecase's service
*/
func (r *resultUsecase) GetService() *service.ResultService {
	return r.service
}

/*
GetRepo get usecase's repo
*/
func (r *resultUsecase) GetRepo() repository.ResultRepository {
	return r.repo
}

func (r *resultUsecase) LoginResult(client *http.Client, studentID, password string) (loginResult model.Result, err error) {
	req := newRequest(studentID, password)
	resp, err := client.Do(req)
	loginResult = r.repo.GetLoginResultByResponse(resp)

	return loginResult, err
}

/*
CurriculumResultBy get curriculum result
@parameter: CurriculumUsecase, string, string
@return: *model.Result, error
*/
func (r *resultUsecase) CurriculumResultBy(curriculumUsecase CurriculumUsecase, studentID, targetStudentID string) (curriculumResult *model.Result, err error) {
	if targetStudentID == "" {
		targetStudentID = studentID
	}

	doc, err := curriculumUsecase.GetService().GetCurriculumDocument(targetStudentID)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	curriculums := curriculumUsecase.GetRepo().ParseCurriculums(doc)

	return r.repo.GetCurriculumResult(curriculums), nil
}

/*
InfoResultBy get info result
@parameter: InfoUsecase, string, string, string, string
@return: *model.Result, err error
*/
func (r *resultUsecase) InfoResultBy(infoUsecase InfoUsecase, studentID, targetStudentID, year, semester string) (curriculumResult *model.Result, err error) {
	if targetStudentID == "" {
		targetStudentID = studentID
	}

	rows, err := infoUsecase.GetService().GetInfoRows(targetStudentID, year, semester)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	info := infoUsecase.GetRepo().GetInfoByRows(rows)

	return r.repo.GetCurriculumCorseResult(info), nil
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
