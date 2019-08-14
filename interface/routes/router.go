package routes

import (
	"tat_gogogo/interface/controllers/courses"
	"tat_gogogo/interface/controllers/curriculum"
	"tat_gogogo/interface/controllers/login"

	"github.com/gin-gonic/gin"
)

/*
RegisterRoutes is a place to register rotes
*/
func RegisterRoutes(router *gin.Engine) {
	router.POST("/login", login.Controller)
	router.POST("/curriculums", curriculum.Controller)
	router.POST("/curriculums/courses", courses.Controller)
}
