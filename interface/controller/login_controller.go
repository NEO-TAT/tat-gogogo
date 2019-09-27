package controller

import (
	"log"
	"tat_gogogo/domain/model"
	"tat_gogogo/usecase"
)

type loginController struct {
	resultUseCase     usecase.ResultUseCase
	curriculumUsecase usecase.CurriculumUseCase
}

/*
LoginController handle login related task
*/
type LoginController interface {
	Login(studentID, password string) (*model.Result, error)
	LoginCurriculum() (bool, error)
}

/*
NewLoginController get a new LoginHandler
*/
func NewLoginController(
	resultUseCase usecase.ResultUseCase,
	curriculumUsecase usecase.CurriculumUseCase,
) LoginController {
	return &loginController{resultUseCase: resultUseCase, curriculumUsecase: curriculumUsecase}
}

/*
Login will login the school system
*/
func (c *loginController) Login(studentID, password string) (*model.Result, error) {
	result, err := c.resultUseCase.LoginResult(studentID, password)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	return result, nil
}

/*
LoginCurriculum will login school curriculum system
*/
func (c *loginController) LoginCurriculum() (bool, error) {
	return c.curriculumUsecase.LoginCurriculum()
}
