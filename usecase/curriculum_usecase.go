package usecase

import (
	"tat_gogogo/domain/model"
	"tat_gogogo/domain/repository"
	"tat_gogogo/domain/service"
)

/*
CurriculumUsecase contains the functions for curriculum usecase
*/
type CurriculumUsecase interface {
	LoginCurriculum() (bool, error)
	IsSameYearAndSem(curriculums []model.Curriculum, year, semester string) bool
	GetService() *service.CurriculumService
	GetRepo() repository.CurriculumRepository
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
GetService get usecase's service
*/
func (c *curriculumUsecase) GetService() *service.CurriculumService {
	return c.service
}

/*
GetRepo get usecase's repo
*/
func (c *curriculumUsecase) GetRepo() repository.CurriculumRepository {
	return c.repo
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
