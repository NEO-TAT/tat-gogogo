package usecase

import (
	"tat_gogogo/domain/model"
	"tat_gogogo/domain/repository"
	"tat_gogogo/domain/service"

	"log"
)

/*
CurriculumUsecase contains the functions for curriculum usecase
*/
type CurriculumUsecase interface {
	GetCurriculums(targetStudentID string) ([]model.Curriculum, error)
	LoginCurriculum() (bool, error)
	IsSameYearAndSem(curriculums []model.Curriculum, year, semester string) bool
}

type curriculumUsecase struct {
	repo    repository.CurriculumRepository
	service *service.CurriculumService
}

/*
NewCurriculumUsecase init a new curriculum usecase
*/
func NewCurriculumUsecase(repo repository.CurriculumRepository, service *service.CurriculumService) CurriculumUsecase {
	return &curriculumUsecase{repo: repo, service: service}
}

/*
LoginCurriculum login curriculum system
*/
func (c *curriculumUsecase) LoginCurriculum() (bool, error) {
	return c.service.IsLoginCurriculum()
}

/*
IsSameYearAndSemBy judge is same year and semester
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
GetCurriculums get []model.Curriculum
*/
func (c *curriculumUsecase) GetCurriculums(targetStudentID string) ([]model.Curriculum, error) {
	doc, err := c.service.GetCurriculumDocument(targetStudentID)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}
	return c.repo.ParseCurriculums(doc), nil
}
