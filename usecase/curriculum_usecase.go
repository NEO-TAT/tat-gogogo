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

func (c *curriculumUsecase) LoginCurriculum() (bool, error) {
	return c.service.IsLoginCurriculum()
}
