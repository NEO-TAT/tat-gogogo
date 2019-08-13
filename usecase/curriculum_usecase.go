package usecase

import (
	"tat_gogogo/domain/repository"
	"tat_gogogo/domain/service"
)

/*
CurriculumUsecase contains the functions for curriculum usecase
*/
type CurriculumUsecase interface {
	LoginCurriculum() (bool, error)
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
func NewCurriculumUsecase(repo repository.CurriculumRepository, service *service.CurriculumService) *curriculumUsecase {
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
