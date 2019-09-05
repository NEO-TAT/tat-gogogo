package usecase

import (
	"tat_gogogo/domain/model"
	"tat_gogogo/domain/repository"
	"tat_gogogo/domain/service"
	"tat_gogogo/utilities/logs"
)

/*
InfoUseCase contains the functions for info usecase
*/
type InfoUseCase interface {
	GetInfo(studentID, year, semester string) (*model.Info, error)
}

type infoUseCase struct {
	repo    repository.InfoRepository
	service *service.InfoService
}

/*
NewInfoUseCase init a new info usecase
*/
func NewInfoUseCase(repo repository.InfoRepository, service *service.InfoService) InfoUseCase {
	return &infoUseCase{repo: repo, service: service}
}

/*
GetInfo get info by studentID, year and semester
*/
func (i *infoUseCase) GetInfo(studentID, year, semester string) (*model.Info, error) {
	rows, err := i.service.GetInfoRows(studentID, year, semester)
	if err != nil {
		logs.Error.Panicln(err)
		return nil, err
	}
	return i.repo.GetInfoByRows(rows), nil
}
