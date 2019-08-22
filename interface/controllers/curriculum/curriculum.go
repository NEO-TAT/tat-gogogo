package curriculum

import (
	"tat_gogogo/domain/model"
	"tat_gogogo/domain/repository"
	"tat_gogogo/domain/service"
	"tat_gogogo/usecase"

	"log"

	"github.com/gin-gonic/gin"
)

type handler struct {
	studentID       string
	password        string
	targetStudentID string
}

/*
Controller is a function for gin to handle curriculum api
*/
func Controller(c *gin.Context) {
	studentID := c.PostForm("studentID")
	password := c.PostForm("password")
	targetStudentID := c.PostForm("targetStudentID")

	handler := newHandler(studentID, password, targetStudentID)

	result, err := handler.login()
	if err != nil {
		c.Status(500)
		return
	}

	if !result.GetSuccess() {
		c.JSON(401, gin.H{
			"message": result.GetData(),
		})
		return
	}

	isLoginCurriculumSuccess, err := handler.loginCurriculum()
	if err != nil {
		c.Status(500)
		return
	}

	if !isLoginCurriculumSuccess {
		c.JSON(401, gin.H{
			"message": "failed to login curriculum",
		})
		return
	}

	curriculumResult, err := handler.getCurriculumResult()
	if err != nil {
		c.Status(500)
		return
	}

	c.JSON(curriculumResult.GetStatus(), gin.H{
		"message": curriculumResult.GetData(),
	})
}

func newHandler(studentID, password, targetStudentID string) *handler {
	return &handler{
		studentID:       studentID,
		password:        password,
		targetStudentID: targetStudentID,
	}
}

func (c *handler) login() (*model.Result, error) {
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

func (c *handler) loginCurriculum() (bool, error) {
	curriculumRepo := repository.NewCurriculumRepository()
	curriculumService := service.NewCurriculumService(curriculumRepo)
	curriculumUsecase := usecase.NewCurriculumUseCase(curriculumRepo, curriculumService)

	return curriculumUsecase.LoginCurriculum()
}

func (c *handler) getCurriculumResult() (*model.Result, error) {
	curriculumResultRepo := repository.NewResultRepository()
	curriculumResultService := service.NewResultService(curriculumResultRepo)
	curriculumResultUsecase := usecase.NewResultUseCase(curriculumResultRepo, curriculumResultService)

	return curriculumResultUsecase.CurriculumResultBy(c.studentID, c.targetStudentID)
}
