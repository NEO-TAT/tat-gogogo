package router

import (
	"tat_gogogo/api/handler"
	"tat_gogogo/api/middleware"
	"tat_gogogo/glob/logs"

	"github.com/gin-gonic/gin"
)

/*
Register is a place to register rotes
*/
func Register(router *gin.Engine) {
	authMiddleware, err := middleware.NewAuthMiddleware()
	if err != nil {
		logs.Error.Panicln(err)
	}

	router.POST("/login", handler.LoginHandler)

	auth := router.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/curriculums/semesters", handler.CurriculumHandler)
		auth.GET("/curriculums/courses", handler.CoursesHandler)
	}
}
