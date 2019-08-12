package curriculum

import (
	"tat_gogogo/domain/repository"
	"tat_gogogo/domain/service"
	"tat_gogogo/usecase"
	"tat_gogogo/utilities/httcli"

	"log"

	"github.com/gin-gonic/gin"
)

/*
Controller is a function for gin to handle curriculum api
@parameter: *gin.Context
*/
func Controller(c *gin.Context) {
	studentID := c.PostForm("studentID")
	password := c.PostForm("password")

	client := httcli.GetInstance()

	loginResultRepo := repository.NewResultRepository()
	loginResultService := service.NewResultService(loginResultRepo)
	loginResultUsecase := usecase.NewResultUsecase(loginResultRepo, loginResultService)

	result, err := loginResultUsecase.Login(client, studentID, password)
	if err != nil {
		log.Panicln(err)
		c.Status(500)
		return
	}

	log.Println("success", result.GetSuccess(), "status", result.GetStatus(), "data", result.GetData())

	if !result.GetSuccess() {
		c.JSON(200, gin.H{"message": "failed to login"})
		return
	}

	curriculumRepo := repository.NewCurriculumRepository()
	curriculumService := service.NewCurriculumService(curriculumRepo)
	curriculumUsecase := usecase.NewCurriculumUsecase(curriculumRepo, curriculumService)

	isLoginCurriculumSuccess, err := curriculumUsecase.LoginCurriculum()
	if err != nil {
		log.Panicln(err)
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{"isLoginCurriculumSuccess": isLoginCurriculumSuccess})
}
