package usecase

import (
	"tat_gogogo/domain/repository"
	"tat_gogogo/domain/service"
)

/*
InfoUsecase contains the functions for info usecase
*/
type InfoUsecase interface {
	GetService() *service.InfoService
	GetRepo() repository.InfoRepository
}

type infoUsecase struct {
	repo    repository.InfoRepository
	service *service.InfoService
}

/*
NewInfoUsecase init a new info usecase
@parameter: repository.CurriculumRepository, *service.CurriculumService
@return: *curriculumUsecase
*/
func NewInfoUsecase(repo repository.InfoRepository, service *service.InfoService) *infoUsecase {
	return &infoUsecase{repo: repo, service: service}
}

/*
GetService get usecase's service
*/
func (r *infoUsecase) GetService() *service.InfoService {
	return r.service
}

/*
GetRepo get usecase's repo
*/
func (r *infoUsecase) GetRepo() repository.InfoRepository {
	return r.repo
}
