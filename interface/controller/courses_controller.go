package controller

import (
	"errors"
	"log"
	"tat_gogogo/domain/model"
	"tat_gogogo/usecase"
)

type coursesController struct {
	resultUseCase     usecase.ResultUseCase
	curriculumUseCase usecase.CurriculumUseCase
}

/*
CoursesController handle courses
*/
type CoursesController interface {
	GetCurriculums(
		studentID,
		targetStudentID string,
	) ([]model.Curriculum, error)
	IsSameYearAndSem(curriculums []model.Curriculum, year, semester string) bool
	GetInfoResult(
		studentID,
		password,
		targetStudentID,
		year,
		semester string,
	) (*model.Result, error)
	GetNoDataResult() *model.Result
}

/*
NewCoursesController get a new CoursesController
*/
func NewCoursesController(
	resultUsecase usecase.ResultUseCase,
	curriculumUseCase usecase.CurriculumUseCase,
) CoursesController {
	return &coursesController{
		resultUseCase:     resultUsecase,
		curriculumUseCase: curriculumUseCase,
	}
}

/*
GetCurriculums get curriculum
*/
func (c *coursesController) GetCurriculums(
	studentID,
	targetStudentID string,
) ([]model.Curriculum, error) {
	curriculumRsult, err := c.resultUseCase.CurriculumResultBy(studentID, targetStudentID)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	if curriculums, ok := curriculumRsult.GetData().([]model.Curriculum); ok {
		return curriculums, nil
	}

	return nil, errors.New("failed to cast []model.Curriculum")
}

/*
IsSameYearAndSem judge is the same year and semester
*/
func (c *coursesController) IsSameYearAndSem(curriculums []model.Curriculum, year, semester string) bool {
	return c.curriculumUseCase.IsSameYearAndSem(curriculums, year, semester)
}

/*
GetInfoResult get info result
*/
func (c *coursesController) GetInfoResult(
	studentID,
	password,
	targetStudentID,
	year,
	semester string,
) (*model.Result, error) {
	return c.resultUseCase.InfoResultBy(studentID, targetStudentID, year, semester)
}

func (c *coursesController) GetNoDataResult() *model.Result {
	return c.resultUseCase.GetNoDataResult()
}
