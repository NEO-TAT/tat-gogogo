package usecase

import (
	"log"
	"tat_gogogo/domain/model"
	"tat_gogogo/domain/repository"
	"tat_gogogo/domain/service"
	"tat_gogogo/utilities/httcli"
)

/*
ResultUseCase contains the functions for result usecase
*/
type ResultUseCase interface {
	LoginResult(studentID, password string) (loginResult *model.Result, err error)
	CurriculumResultBy(studentID, targetStudentID string) (curriculumResult *model.Result, err error)
	InfoResultBy(studentID, targetStudentID, year, semester string) (curriculumResult *model.Result, err error)
	GetNoDataResult() *model.Result
}

type resultUsecase struct {
	resultRepo    	repository.ResultRepository
	curriculumRepo 	repository.CurriculumRepository
	infoRepo 				repository.InfoRepository
	service 				*service.ResultService
}

/*
NewResultUseCase init a new result usecase
*/
func NewResultUseCase(
	resultRepo repository.ResultRepository, 
	curriculumRepo repository.CurriculumRepository, 
	infoRepo repository.InfoRepository,
	service *service.ResultService,
	) ResultUseCase {
	return &resultUsecase{
		resultRepo: resultRepo,
		curriculumRepo: curriculumRepo,
		infoRepo: infoRepo,
		service: service,
	}
}

func (r *resultUsecase) LoginResult(studentID, password string) (loginResult *model.Result, err error) {
	req := r.service.NewLoginRequest(studentID, password)
	client := httcli.GetInstance()
	resp, err := client.Do(req)
	loginResult = r.resultRepo.GetLoginResultByResponse(resp)

	return loginResult, err
}

/*
CurriculumResultBy get curriculum result
*/
func (r *resultUsecase) CurriculumResultBy(studentID, targetStudentID string) (curriculumResult *model.Result, err error) {
	if targetStudentID == "" {
		targetStudentID = studentID
	}

	curriculumService := service.NewCurriculumService(r.curriculumRepo)
	curriculumUsecase := NewCurriculumUseCase(r.curriculumRepo, curriculumService)

	curriculums, err := curriculumUsecase.GetCurriculums(targetStudentID)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	return r.resultRepo.GetCurriculumResult(curriculums), nil
}

/*
InfoResultBy get info result
*/
func (r *resultUsecase) InfoResultBy(studentID, targetStudentID, year, semester string) (curriculumResult *model.Result, err error) {
	if targetStudentID == "" {
		targetStudentID = studentID
	}

	infoService := service.NewInfoService(r.infoRepo)
	infoUsecase := NewInfoUseCase(r.infoRepo, infoService)

	info, err := infoUsecase.GetInfo(targetStudentID, year, semester)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	return r.resultRepo.GetCurriculumCorseResult(info), nil
}

/*
GetNoDataResult get no data result
*/
func (r *resultUsecase) GetNoDataResult() *model.Result {
	return model.NewResult(false, 400, "查無該學年或學期資料")
}
