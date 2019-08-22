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

type handler struct {
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

	handler := newHandler(studentID, password, targetStudentID, year, semester)

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
			"message": "登入課程系統失敗",
		})
		return
	}

	curriculums, err := handler.getCurriculums()
	if err != nil {
		c.Status(500)
		return
	}

	isSameYearAndSem := handler.isSameYearAndSem(curriculums)

	if !isSameYearAndSem {
		result := getNoDataResult()
		c.JSON(result.GetStatus(), gin.H{
			"message": result.GetData(),
		})
		return
	}

	infoResult, err := handler.getInfoResult()
	if err != nil {
		log.Panicln(err)
		c.Status(500)
		return
	}

	c.JSON(infoResult.GetStatus(), infoResult.GetData())

}

func newHandler(studentID, password, targetStudentID, year, semester string) *handler {
	return &handler{
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
	resultUsecase := usecase.NewResultUseCase(resultRepo, resultService)
	return resultUsecase.GetNoDataResult()
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

func (c *handler) getCurriculums() ([]model.Curriculum, error) {
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

func (c *handler) isSameYearAndSem(curriculums []model.Curriculum) bool {
	curriculumRepo := repository.NewCurriculumRepository()
	curriculumService := service.NewCurriculumService(curriculumRepo)
	curriculumUsecase := usecase.NewCurriculumUseCase(curriculumRepo, curriculumService)

	return curriculumUsecase.IsSameYearAndSem(curriculums, c.year, c.semester)
}

func (c *handler) getInfoResult() (*model.Result, error) {
	infoResultRepo := repository.NewResultRepository()
	infoResultService := service.NewResultService(infoResultRepo)
	infoResultUsecase := usecase.NewResultUseCase(infoResultRepo, infoResultService)

	return infoResultUsecase.InfoResultBy(c.studentID, c.targetStudentID, c.year, c.semester)
}
