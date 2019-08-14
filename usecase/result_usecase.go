package usecase

import (
	"log"
	"net/http"
	"tat_gogogo/domain/model"
	"tat_gogogo/domain/repository"
	"tat_gogogo/domain/service"
	"tat_gogogo/utilities/httcli"
)

/*
ResultUsecase contains the functions for result usecase
*/
type ResultUsecase interface {
	LoginResult(client *http.Client, studentID, password string) (loginResult *model.Result, err error)
	CurriculumResultBy(curriculumUsecase CurriculumUsecase, studentID, targetStudentID, year, semester string) (curriculumResult model.Result, err error)
	InfoResultBy(infoUsecase InfoUsecase, studentID, targetStudentID, year, semester string) (curriculumResult model.Result, err error)
	GetNoDataResult() *model.Result
}

type resultUsecase struct {
	repo    repository.ResultRepository
	service *service.ResultService
}

/*
NewResultUsecase init a new result usecase
*/
func NewResultUsecase(repo repository.ResultRepository, service *service.ResultService) *resultUsecase {
	return &resultUsecase{repo: repo, service: service}
}

func (r *resultUsecase) LoginResult(studentID, password string) (loginResult *model.Result, err error) {
	req := r.service.NewLoginRequest(studentID, password)
	client := httcli.GetInstance()
	resp, err := client.Do(req)
	loginResult = r.repo.GetLoginResultByResponse(resp)

	return loginResult, err
}

/*
CurriculumResultBy get curriculum result
*/
func (r *resultUsecase) CurriculumResultBy(studentID, targetStudentID string) (curriculumResult *model.Result, err error) {
	if targetStudentID == "" {
		targetStudentID = studentID
	}

	curriculumRepo := repository.NewCurriculumRepository()
	curriculumService := service.NewCurriculumService(curriculumRepo)
	curriculumUsecase := NewCurriculumUsecase(curriculumRepo, curriculumService)

	curriculums, err := curriculumUsecase.GetCurriculums(targetStudentID)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	return r.repo.GetCurriculumResult(curriculums), nil
}

/*
InfoResultBy get info result
*/
func (r *resultUsecase) InfoResultBy(studentID, targetStudentID, year, semester string) (curriculumResult *model.Result, err error) {
	if targetStudentID == "" {
		targetStudentID = studentID
	}

	infoRepo := repository.NewInfoRepository()
	infoService := service.NewInfoService(infoRepo)
	infoUsecase := NewInfoUsecase(infoRepo, infoService)

	info, err := infoUsecase.GetInfo(targetStudentID, year, semester)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	return r.repo.GetCurriculumCorseResult(info), nil
}

/*
GetNoDataResult get no data result
*/
func (r *resultUsecase) GetNoDataResult() *model.Result {
	return model.NewResult(false, 400, "查無該學年或學期資料")
}
