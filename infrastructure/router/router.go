package router

import (
	"tat_gogogo/infrastructure/api/handler"
	"tat_gogogo/infrastructure/middleware"
	"tat_gogogo/utilities/logs"

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

	// -----------------------------------------------[API List]

	router.POST("/login", handler.LoginHandler)

	auth := router.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/curriculums/semesters", handler.CurriculumHandler)
		auth.GET("/curriculums/courses", handler.CoursesHandler)
	}
}
