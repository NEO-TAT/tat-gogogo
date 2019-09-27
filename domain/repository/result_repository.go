package repository

import (
	"net/http"
	"tat_gogogo/domain/model"
)

/*
ResultRepository declares repo of result
*/
type ResultRepository interface {
	GetLoginResultByResponse(resp *http.Response) *model.Result
	GetCurriculumResult(cirriculums []model.Curriculum) *model.Result
	GetCurriculumCorseResult(info *model.Info) *model.Result
	GetNoDataResult() *model.Result
}
