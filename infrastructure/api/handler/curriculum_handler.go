package handler

import (
	"tat_gogogo/domain/model"
	"tat_gogogo/domain/repository"
	"tat_gogogo/domain/service"
	"tat_gogogo/usecase"
)

type curriculumHandler struct {
	studentID       string
	password        string
	targetStudentID string
}

/*
CurriculumHandler handle curriculum flow
*/
type CurriculumHandler interface {
	GetCurriculumResult() (*model.Result, error)
}

/*
NewCurriculumHandler get a new CurriculumHandler
*/
func NewCurriculumHandler(studentID, password, targetStudentID string) CurriculumHandler {
	return &curriculumHandler{
		studentID:       studentID,
		password:        password,
		targetStudentID: targetStudentID,
	}
}

/*
GetCurriculumResult get curriculum
*/
func (c *curriculumHandler) GetCurriculumResult() (*model.Result, error) {
	curriculumResultRepo := repository.NewResultRepository()
	curriculumResultService := service.NewResultService(curriculumResultRepo)
	curriculumResultUsecase := usecase.NewResultUseCase(curriculumResultRepo, curriculumResultService)

	return curriculumResultUsecase.CurriculumResultBy(c.studentID, c.targetStudentID)
}
