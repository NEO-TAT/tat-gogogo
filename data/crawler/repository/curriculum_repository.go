package repository

import (
	"strings"
	"tat_gogogo/domain/model"
	"tat_gogogo/domain/repository"

	"github.com/PuerkitoBio/goquery"
)

type curriculumRepository struct{}

/*
NewCurriculumRepository init a curriculumRepository
*/
func NewCurriculumRepository() repository.CurriculumRepository {
	return &curriculumRepository{}
}

/*
ParseCurriculums parse the curriculum from doc
*/
func (c *curriculumRepository) ParseCurriculums(doc *goquery.Document) []model.Curriculum {
	var curriculums []model.Curriculum

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		if href, ok := s.Attr("href"); ok {
			splits := strings.Split(href, "&")
			year := strings.Replace(splits[2], "year=", "", 1)
			sem := strings.Replace(splits[3], "sem=", "", 1)

			curriculum := model.Curriculum{Year: year, Semester: sem}
			curriculums = append(curriculums, curriculum)
		}
	})

	return curriculums
}
