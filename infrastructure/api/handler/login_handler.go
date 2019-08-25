package handler

import (
	"log"
	"tat_gogogo/domain/model"
	"tat_gogogo/domain/repository"
	"tat_gogogo/domain/service"
	"tat_gogogo/usecase"
)

type loginHandler struct {
	studentID string
	password  string
}

/*
LoginHandler handle login related task
*/
type LoginHandler interface {
	Login() (*model.Result, error)
	LoginCurriculum() (bool, error)
}

/*
NewLoginHanlder get a new LoginHandler
*/
func NewLoginHanlder(studentID, password string) LoginHandler {
	return &loginHandler{
		studentID: studentID,
		password:  password,
	}
}

/*
Login will login the school system
*/
func (c *loginHandler) Login() (*model.Result, error) {
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
func (c *loginHandler) LoginCurriculum() (bool, error) {
	curriculumRepo := repository.NewCurriculumRepository()
	curriculumService := service.NewCurriculumService(curriculumRepo)
	curriculumUsecase := usecase.NewCurriculumUseCase(curriculumRepo, curriculumService)

	return curriculumUsecase.LoginCurriculum()
}
