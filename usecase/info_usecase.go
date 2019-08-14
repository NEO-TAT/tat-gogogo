package usecase

import (
	"tat_gogogo/domain/model"
	"tat_gogogo/domain/repository"
	"tat_gogogo/domain/service"

	"github.com/PuerkitoBio/goquery"
)

/*
InfoUsecase contains the functions for info usecase
*/
type InfoUsecase interface {
	GetInfoRows(studentID, year, semester string) (*goquery.Selection, error)
	GetInfoByRows(rows *goquery.Selection) *model.Info
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
func NewInfoUsecase(repo repository.InfoRepository, service *service.InfoService) InfoUsecase {
	return &infoUsecase{repo: repo, service: service}
}

/*
GetInfoRows get ifno rows
@parameter: string, string, string
@return: *goquery.Selection, error
*/
func (i *infoUsecase) GetInfoRows(studentID, year, semester string) (*goquery.Selection, error) {
	return i.service.GetInfoRows(studentID, year, semester)
}

/*
GetInfoByRows get info by selection
@parameter: *goquery.Selection
@return: *model.Info
*/
func (i *infoUsecase) GetInfoByRows(rows *goquery.Selection) *model.Info {
	return i.repo.GetInfoByRows(rows)
}