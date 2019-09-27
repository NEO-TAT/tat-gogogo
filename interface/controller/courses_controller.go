package controller

import (
	"errors"
	"log"
	"tat_gogogo/data/crawler/repository"
	"tat_gogogo/domain/model"
	"tat_gogogo/domain/service"
	"tat_gogogo/usecase"
)

type coursesController struct {
	studentID       string
	password        string
	targetStudentID string
	year            string
	semester        string
}

/*
CoursesController handle courses
*/
type CoursesController interface {
	GetCurriculums() ([]model.Curriculum, error)
	IsSameYearAndSem(curriculums []model.Curriculum) bool
	GetInfoResult() (*model.Result, error)
	GetNoDataResult() *model.Result
}

/*
NewCoursesController get a new CoursesController
*/
func NewCoursesController(studentID, password, targetStudentID, year, semester string) CoursesController {
	return &coursesController{
		studentID:       studentID,
		password:        password,
		targetStudentID: targetStudentID,
		year:            year,
		semester:        semester,
	}
}

/*
GetCurriculums get curriculum
*/
func (c *coursesController) GetCurriculums() ([]model.Curriculum, error) {
	curriculumResultRepo := repository.NewResultRepository()
	curriculumRepo := repository.NewCurriculumRepository()
	infoRepo := repository.NewInfoRepository()

	curriculumResultService := service.NewResultService(curriculumResultRepo)
	curriculumResultUsecase := usecase.NewResultUseCase(
		curriculumResultRepo,
		curriculumRepo,
		infoRepo,
		curriculumResultService,
	)

	curriculumRsult, err := curriculumResultUsecase.CurriculumResultBy(c.studentID, c.targetStudentID)
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
func (c *coursesController) IsSameYearAndSem(curriculums []model.Curriculum) bool {
	curriculumRepo := repository.NewCurriculumRepository()
	curriculumService := service.NewCurriculumService(curriculumRepo)
	curriculumUsecase := usecase.NewCurriculumUseCase(curriculumRepo, curriculumService)

	return curriculumUsecase.IsSameYearAndSem(curriculums, c.year, c.semester)
}

/*
GetInfoResult get info result
*/
func (c *coursesController) GetInfoResult() (*model.Result, error) {
	resultRepo := repository.NewResultRepository()
	curriculumRepo := repository.NewCurriculumRepository()
	infoRepo := repository.NewInfoRepository()

	infoResultService := service.NewResultService(resultRepo)
	infoResultUsecase := usecase.NewResultUseCase(
		resultRepo,
		curriculumRepo,
		infoRepo,
		infoResultService,
	)

	return infoResultUsecase.InfoResultBy(c.studentID, c.targetStudentID, c.year, c.semester)
}

func (c *coursesController) GetNoDataResult() *model.Result {
	resultRepo := repository.NewResultRepository()
	curriculumRepo := repository.NewCurriculumRepository()
	infoRepo := repository.NewInfoRepository()

	resultService := service.NewResultService(resultRepo)
	resultUsecase := usecase.NewResultUseCase(
		resultRepo,
		curriculumRepo,
		infoRepo,
		resultService)
	return resultUsecase.GetNoDataResult()
}
