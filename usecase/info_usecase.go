package usecase

import (
	"log"
	"tat_gogogo/data/crawler/service"
	"tat_gogogo/domain/model"
	"tat_gogogo/domain/repository"
)

/*
InfoUseCase contains the functions for info usecase
*/
type InfoUseCase interface {
	GetInfo(studentID, year, semester string) (*model.Info, error)
}

type infoUseCase struct {
	infoRepo repository.InfoRepository
	service  *service.InfoService
}

/*
NewInfoUseCase init a new info usecase
*/
func NewInfoUseCase(infoRepo repository.InfoRepository, service *service.InfoService) InfoUseCase {
	return &infoUseCase{infoRepo: infoRepo, service: service}
}

/*
GetInfo get info by studentID, year and semester
*/
func (i *infoUseCase) GetInfo(studentID, year, semester string) (*model.Info, error) {
	rows, err := i.service.GetInfoRows(studentID, year, semester)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}
	return i.infoRepo.GetInfoByRows(rows), nil
}
