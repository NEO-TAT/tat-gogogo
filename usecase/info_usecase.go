package usecase

import (
	"log"
	"tat_gogogo/domain/model"
	"tat_gogogo/domain/repository"
	"tat_gogogo/domain/service"
)

/*
InfoUsecase contains the functions for info usecase
*/
type InfoUsecase interface {
	GetInfo(studentID, year, semester string) (*model.Info, error)
}

type infoUsecase struct {
	repo    repository.InfoRepository
	service *service.InfoService
}

/*
NewInfoUsecase init a new info usecase
*/
func NewInfoUsecase(repo repository.InfoRepository, service *service.InfoService) InfoUsecase {
	return &infoUsecase{repo: repo, service: service}
}

/*
GetInfo get info by studentID, year and semester
*/
func (i *infoUsecase) GetInfo(studentID, year, semester string) (*model.Info, error) {
	rows, err := i.service.GetInfoRows(studentID, year, semester)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}
	return i.repo.GetInfoByRows(rows), nil
}
