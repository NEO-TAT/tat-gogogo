package controller

import (
	"tat_gogogo/data/crawler/repository"
	"tat_gogogo/domain/model"
	"tat_gogogo/domain/service"
	"tat_gogogo/usecase"
)

type curriculumController struct {
	studentID       string
	password        string
	targetStudentID string
}

/*
CurriculumController handle curriculum flow
*/
type CurriculumController interface {
	GetCurriculumResult() (*model.Result, error)
}

/*
NewCurriculumController get a new CurriculumController
*/
func NewCurriculumController(studentID, password, targetStudentID string) CurriculumController {
	return &curriculumController{
		studentID:       studentID,
		password:        password,
		targetStudentID: targetStudentID,
	}
}

/*
GetCurriculumResult get curriculum
*/
func (c *curriculumController) GetCurriculumResult() (*model.Result, error) {
	resultRepo := repository.NewResultRepository()
	curriculumRepo := repository.NewCurriculumRepository()
	infoRepo := repository.NewInfoRepository()

	curriculumResultService := service.NewResultService(resultRepo)
	curriculumResultUsecase := usecase.NewResultUseCase(
		resultRepo,
		curriculumRepo,
		infoRepo,
		curriculumResultService,
	)

	return curriculumResultUsecase.CurriculumResultBy(c.studentID, c.targetStudentID)
}
