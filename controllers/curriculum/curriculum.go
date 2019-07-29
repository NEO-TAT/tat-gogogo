package curriculum

import (
	"tat_gogogo/crawler/curriculum"

	"log"

	"github.com/gin-gonic/gin"
)

/*
Controller handle curriculum login api
*/
func Controller(c *gin.Context) {
	studentID := c.PostForm("studentId")
	password := c.PostForm("password")
	targetStudentID := c.PostForm("targetStudentId")

	result, err := curriculum.GetCurriculums(studentID, password, targetStudentID)
	if err != nil {
		log.Panicln(err)
		c.Status(500)
		return
	}

	if result.Status != 200 {
		c.JSON(result.Status, gin.H{
			"message": result.Data,
		})
		return
	}

	c.JSON(result.Status, result.Data)
}

/*
CourseController handle search course
the default target student will be self
*/
func CourseController(c *gin.Context) {
	studentID := c.PostForm("studentId")
	password := c.PostForm("password")
	targetStudentID := c.PostForm("targetStudentId")
	year := c.PostForm("year")
	sem := c.PostForm("sem")

	result, err := curriculum.GetCurriculumCourse(studentID, password, targetStudentID, year, sem)
	if err != nil {
		log.Panicln(err)
		c.Status(500)
		return
	}

	if result.Status != 200 {
		c.JSON(result.Status, gin.H{
			"message": result.Data,
		})
		return
	}

	c.JSON(result.Status, result.Data)
}
