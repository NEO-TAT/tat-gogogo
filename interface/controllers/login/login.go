package login

import (
	"tat_gogogo/domain/repository"
	"tat_gogogo/domain/service"
	"tat_gogogo/usecase"
	"tat_gogogo/utilities/httcli"

	"log"

	"github.com/gin-gonic/gin"
)

/*
Controller is a function for gin to handle login api
@parameter: *gin.Context
*/
func Controller(c *gin.Context) {
	studentID := c.PostForm("studentID")
	password := c.PostForm("password")
	client := httcli.GetInstance()

	repo := repository.NewResultRepository()
	service := service.NewResultService(repo)
	resultUsecase := usecase.NewResultUsecase(repo, service)
	result, err := resultUsecase.LoginResult(client, studentID, password)

	if err != nil {
		log.Panicln("failed to fetch login cookie")
		c.Status(500)
		return
	}

	if result.GetStatus() != 200 {
		c.JSON(result.GetStatus(), gin.H{
			"message": result.GetData(),
		})
		return
	}

	c.Status(200)

}
