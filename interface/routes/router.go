package routes

import (
	"log"
	"tat_gogogo/interface/controllers/courses"
	"tat_gogogo/interface/controllers/curriculum"
	"tat_gogogo/interface/controllers/login"
	"tat_gogogo/interface/jwt"

	"github.com/gin-gonic/gin"
)

/*
RegisterRoutes is a place to register rotes
*/
func RegisterRoutes(router *gin.Engine) {
	authMiddleware, err := jwt.AuthMiddleware()
	if err != nil {
		log.Panicln(err)
	}

	router.POST("/login", login.Controller)

	auth := router.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/curriculums/semesters", curriculum.Controller)
		auth.GET("/curriculums/courses", courses.Controller)
	}
}
