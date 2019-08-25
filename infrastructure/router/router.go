package router

import (
	"log"
	"tat_gogogo/infrastructure/api/handler"
	"tat_gogogo/infrastructure/middleware"

	"github.com/gin-gonic/gin"
)

/*
Register is a place to register rotes
*/
func Register(router *gin.Engine) {
	authMiddleware, err := middleware.NewAuthMiddleware()
	if err != nil {
		log.Panicln(err)
	}

	router.POST("/login", handler.LoginHandler)

	auth := router.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/curriculums/semesters", handler.CurriculumHandler)
		auth.GET("/curriculums/courses", handler.CoursesHandler)
	}
}
