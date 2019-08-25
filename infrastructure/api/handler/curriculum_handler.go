package handler

import (
	"log"
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
	Login() (*model.Result, error)
	LoginCurriculum() (bool, error)
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
Login will login the school system
*/
func (c *curriculumHandler) Login() (*model.Result, error) {
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
func (c *curriculumHandler) LoginCurriculum() (bool, error) {
	curriculumRepo := repository.NewCurriculumRepository()
	curriculumService := service.NewCurriculumService(curriculumRepo)
	curriculumUsecase := usecase.NewCurriculumUseCase(curriculumRepo, curriculumService)

	return curriculumUsecase.LoginCurriculum()
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
