package handler

import (
	"tat_gogogo/interface/controller"

	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/gin-gonic/gin"
)

/*
CurriculumHandler is a function for gin to handle curriculum api
*/
func CurriculumHandler(c *gin.Context) {
	targetStudentID := c.Query("targetStudentID")

	claims := jwt.ExtractClaims(c)
	studentID := claims["studentID"].(string)
	password := claims["password"].(string)

	loginController := controller.NewLoginController(studentID, password)
	curriculumController := controller.NewCurriculumController(studentID, password, targetStudentID)

	result, err := loginController.Login()
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

	isLoginCurriculumSuccess, err := loginController.LoginCurriculum()
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

	curriculumResult, err := curriculumController.GetCurriculumResult()
	if err != nil {
		c.Status(500)
		return
	}

	c.JSON(curriculumResult.GetStatus(), curriculumResult.GetData())
}
