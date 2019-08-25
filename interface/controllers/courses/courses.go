package courses

import (
	"log"
	"tat_gogogo/domain/model"
	"tat_gogogo/domain/repository"
	"tat_gogogo/domain/service"
	"tat_gogogo/infrastructure/api/handler"
	"tat_gogogo/usecase"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

/*
Controller is a function for gin to handle courses api
*/
func Controller(c *gin.Context) {
	targetStudentID := c.Query("targetStudentID")
	year := c.Query("year")
	semester := c.Query("semester")

	claims := jwt.ExtractClaims(c)
	studentID := claims["studentID"].(string)
	password := claims["password"].(string)

	loginHandler := handler.NewLoginHanlder(studentID, password)
	courseHandler := handler.NewCoursesHandler(studentID, password, targetStudentID, year, semester)

	result, err := loginHandler.Login()
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

	isLoginCurriculumSuccess, err := loginHandler.LoginCurriculum()
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

	curriculums, err := courseHandler.GetCurriculums()
	if err != nil {
		c.Status(500)
		return
	}

	isSameYearAndSem := courseHandler.IsSameYearAndSem(curriculums)

	if !isSameYearAndSem {
		result := getNoDataResult()
		c.JSON(result.GetStatus(), gin.H{
			"message": result.GetData(),
		})
		return
	}

	infoResult, err := courseHandler.GetInfoResult()
	if err != nil {
		log.Panicln(err)
		c.Status(500)
		return
	}

	c.JSON(infoResult.GetStatus(), infoResult.GetData())

}

func getNoDataResult() *model.Result {
	resultRepo := repository.NewResultRepository()
	resultService := service.NewResultService(resultRepo)
	resultUsecase := usecase.NewResultUseCase(resultRepo, resultService)
	return resultUsecase.GetNoDataResult()
}
