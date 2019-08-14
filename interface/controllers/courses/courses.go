package courses

import (
	"log"
	"tat_gogogo/domain/model"
	"tat_gogogo/domain/repository"
	"tat_gogogo/domain/service"
	"tat_gogogo/usecase"

	"errors"

	"github.com/gin-gonic/gin"
)

type controller struct {
	studentID       string
	password        string
	targetStudentID string
	year            string
	semester        string
}

/*
Controller is a function for gin to handle courses api
*/
func Controller(c *gin.Context) {
	studentID := c.PostForm("studentID")
	password := c.PostForm("password")
	targetStudentID := c.PostForm("targetStudentID")
	year := c.PostForm("year")
	semester := c.PostForm("semester")

	controller := newController(studentID, password, targetStudentID, year, semester)

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

	curriculums, err := controller.getCurriculums()
	if err != nil {
		c.Status(500)
		return
	}

	isSameYearAndSem := controller.isSameYearAndSem(curriculums)

	if !isSameYearAndSem {
		result := getNoDataResult()
		c.JSON(result.GetStatus(), gin.H{
			"message": result.GetData(),
		})
		return
	}

	infoResult, err := controller.getInfoResult()
	if err != nil {
		log.Panicln(err)
		c.Status(500)
		return
	}

	c.JSON(infoResult.GetStatus(), gin.H{
		"message": infoResult.GetData(),
	})

}

func newController(studentID, password, targetStudentID, year, semester string) *controller {
	return &controller{
		studentID:       studentID,
		password:        password,
		targetStudentID: targetStudentID,
		year:            year,
		semester:        semester,
	}
}

func getNoDataResult() *model.Result {
	resultRepo := repository.NewResultRepository()
	resultService := service.NewResultService(resultRepo)
	resultUsecase := usecase.NewResultUsecase(resultRepo, resultService)
	return resultUsecase.GetNoDataResult()
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

func (c *controller) getCurriculums() ([]model.Curriculum, error) {
	curriculumResultRepo := repository.NewResultRepository()
	curriculumResultService := service.NewResultService(curriculumResultRepo)
	curriculumResultUsecase := usecase.NewResultUsecase(curriculumResultRepo, curriculumResultService)

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

func (c *controller) isSameYearAndSem(curriculums []model.Curriculum) bool {
	curriculumRepo := repository.NewCurriculumRepository()
	curriculumService := service.NewCurriculumService(curriculumRepo)
	curriculumUsecase := usecase.NewCurriculumUsecase(curriculumRepo, curriculumService)

	return curriculumUsecase.IsSameYearAndSem(curriculums, c.year, c.semester)
}

func (c *controller) getInfoResult() (*model.Result, error) {
	infoResultRepo := repository.NewResultRepository()
	infoResultService := service.NewResultService(infoResultRepo)
	infoResultUsecase := usecase.NewResultUsecase(infoResultRepo, infoResultService)

	return infoResultUsecase.InfoResultBy(c.studentID, c.targetStudentID, c.year, c.semester)
}
