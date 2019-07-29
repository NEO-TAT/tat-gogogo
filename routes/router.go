package routes

import (
	"tat_gogogo/controllers/curriculum"
	"tat_gogogo/controllers/login"

	"github.com/gin-gonic/gin"
)

/*
RegisterRoutes is a place to register rotes
*/
func RegisterRoutes(router *gin.Engine) {

	router.POST("/login", login.Controller)
	router.POST("/curriculums", curriculum.Controller)
	router.POST("/curriculums/course", curriculum.CourseController)
}
