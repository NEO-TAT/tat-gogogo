package handler

import (
	"errors"
	"log"
	"tat_gogogo/domain/model"
	"tat_gogogo/domain/repository"
	"tat_gogogo/domain/service"
	"tat_gogogo/usecase"
)

type coursesHandler struct {
	studentID       string
	password        string
	targetStudentID string
	year            string
	semester        string
}

/*
CoursesHandler handle courses
*/
type CoursesHandler interface {
	Login() (*model.Result, error)
	LoginCurriculum() (bool, error)
	GetCurriculums() ([]model.Curriculum, error)
	IsSameYearAndSem(curriculums []model.Curriculum) bool
	GetInfoResult() (*model.Result, error)
}

/*
NewCoursesHandler get a new CoursesHandler
*/
func NewCoursesHandler(studentID, password, targetStudentID, year, semester string) CoursesHandler {
	return &coursesHandler{
		studentID:       studentID,
		password:        password,
		targetStudentID: targetStudentID,
		year:            year,
		semester:        semester,
	}
}

/*
Login will login the school system
*/
func (c *coursesHandler) Login() (*model.Result, error) {
	loginResultRepo := repository.NewResultRepository()
	loginResultService := service.NewResultService(loginResultRepo)
	loginResultUsecase := usecase.NewResultUseCase(loginResultRepo, loginResultService)

	result, err := loginResultUsecase.LoginResult(c.studentID, c.password)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	return result, nil
}

/*
LoginCurriculum will login school curriculum system
*/
func (c *coursesHandler) LoginCurriculum() (bool, error) {
	curriculumRepo := repository.NewCurriculumRepository()
	curriculumService := service.NewCurriculumService(curriculumRepo)
	curriculumUsecase := usecase.NewCurriculumUseCase(curriculumRepo, curriculumService)

	return curriculumUsecase.LoginCurriculum()
}

/*
GetCurriculums get curriculum
*/
func (c *coursesHandler) GetCurriculums() ([]model.Curriculum, error) {
	curriculumResultRepo := repository.NewResultRepository()
	curriculumResultService := service.NewResultService(curriculumResultRepo)
	curriculumResultUsecase := usecase.NewResultUseCase(curriculumResultRepo, curriculumResultService)

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
func (c *coursesHandler) IsSameYearAndSem(curriculums []model.Curriculum) bool {
	curriculumRepo := repository.NewCurriculumRepository()
	curriculumService := service.NewCurriculumService(curriculumRepo)
	curriculumUsecase := usecase.NewCurriculumUseCase(curriculumRepo, curriculumService)

	return curriculumUsecase.IsSameYearAndSem(curriculums, c.year, c.semester)
}

/*
GetInfoResult get info result
*/
func (c *coursesHandler) GetInfoResult() (*model.Result, error) {
	infoResultRepo := repository.NewResultRepository()
	infoResultService := service.NewResultService(infoResultRepo)
	infoResultUsecase := usecase.NewResultUseCase(infoResultRepo, infoResultService)

	return infoResultUsecase.InfoResultBy(c.studentID, c.targetStudentID, c.year, c.semester)
}
