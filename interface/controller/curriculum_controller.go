package controller

import (
	"tat_gogogo/domain/model"
	"tat_gogogo/usecase"
)

type curriculumController struct {
	resultUseCase usecase.ResultUseCase
}

/*
CurriculumController handle curriculum flow
*/
type CurriculumController interface {
	GetCurriculumResult(studentID, targetStudentID string) (*model.Result, error)
}

/*
NewCurriculumController get a new CurriculumController
*/
func NewCurriculumController(resultUseCase usecase.ResultUseCase) CurriculumController {
	return &curriculumController{
		resultUseCase: resultUseCase,
	}
}

/*
GetCurriculumResult get curriculum
*/
func (c *curriculumController) GetCurriculumResult(studentID, targetStudentID string) (*model.Result, error) {
	return c.resultUseCase.CurriculumResultBy(studentID, targetStudentID)
}
