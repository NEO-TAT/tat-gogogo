package curriculum

import (
	"tat_gogogo/domain/model"
	"tat_gogogo/domain/repository"
	"tat_gogogo/domain/service"
	"tat_gogogo/usecase"

	"log"

	"github.com/gin-gonic/gin"
)

type controller struct {
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

	controller := newController(studentID, password, targetStudentID)

	result, err := controller.login()
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

	isLoginCurriculumSuccess, err := controller.loginCurriculum()
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

	curriculumResult, err := controller.getCurriculumResult()
	if err != nil {
		c.Status(500)
		return
	}

	c.JSON(curriculumResult.GetStatus(), gin.H{
		"message": curriculumResult.GetData(),
	})
}

func newController(studentID, password, targetStudentID string) *controller {
	return &controller{
		studentID:       studentID,
		password:        password,
		targetStudentID: targetStudentID,
	}
}

func (c *controller) login() (*model.Result, error) {
	loginResultRepo := repository.NewResultRepository()
	loginResultService := service.NewResultService(loginResultRepo)
	loginResultUsecase := usecase.NewResultUsecase(loginResultRepo, loginResultService)

	result, err := loginResultUsecase.LoginResult(c.studentID, c.password)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	return result, nil
}

func (c *controller) loginCurriculum() (bool, error) {
	curriculumRepo := repository.NewCurriculumRepository()
	curriculumService := service.NewCurriculumService(curriculumRepo)
	curriculumUsecase := usecase.NewCurriculumUsecase(curriculumRepo, curriculumService)

	return curriculumUsecase.LoginCurriculum()
}

func (c *controller) getCurriculumResult() (*model.Result, error) {
	curriculumResultRepo := repository.NewResultRepository()
	curriculumResultService := service.NewResultService(curriculumResultRepo)
	curriculumResultUsecase := usecase.NewResultUsecase(curriculumResultRepo, curriculumResultService)

	return curriculumResultUsecase.CurriculumResultBy(c.studentID, c.targetStudentID)
}
