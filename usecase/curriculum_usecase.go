package usecase

import (
	"strings"
	"tat_gogogo/domain/model"
	"tat_gogogo/domain/repository"
	"tat_gogogo/domain/service"

	"github.com/PuerkitoBio/goquery"
)

/*
CurriculumUsecase contains the functions for curriculum usecase
*/
type CurriculumUsecase interface {
	LoginCurriculum() (bool, error)
	IsSameYearAndSem(curriculums []model.Curriculum, year, semester string) bool
	GetCurriculumDocument(targetStudentID string) (*goquery.Document, error)
	ParseCurriculums(doc *goquery.Document) []model.Curriculum
}

type curriculumUsecase struct {
	repo    repository.CurriculumRepository
	service *service.CurriculumService
}

/*
NewCurriculumUsecase init a new curriculum usecase
@parameter: repository.CurriculumRepository, *service.CurriculumService
@return: *curriculumUsecase
*/
func NewCurriculumUsecase(repo repository.CurriculumRepository, service *service.CurriculumService) CurriculumUsecase {
	return &curriculumUsecase{repo: repo, service: service}
}

/*
LoginCurriculum login curriculum system
@return bool, error
*/
func (c *curriculumUsecase) LoginCurriculum() (bool, error) {
	return c.service.IsLoginCurriculum()
}

/*
IsSameYearAndSemBy judge is same year and semester
@parameter: []model.Curriculum, string, string
@return: bool
*/
func (c *curriculumUsecase) IsSameYearAndSem(curriculums []model.Curriculum, year, semester string) bool {
	for _, curriculum := range curriculums {
		if curriculum.Year == year && curriculum.Semester == semester {
			return true
		}
	}
	return false
}

/*
GetCurriculumDocument will get curriculum doc from the NewRequest
@paramter: targetStudentID string
@return: *goquery.Document, error
*/
func (c *curriculumUsecase) GetCurriculumDocument(targetStudentID string) (*goquery.Document, error) {
	return c.service.GetCurriculumDocument(targetStudentID)
}

/*
ParseCurriculums parse the curriculum from doc
@parameter: *goquery.Document
@return: []model.Curriculum
*/
func (c *curriculumUsecase) ParseCurriculums(doc *goquery.Document) []model.Curriculum {
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
