package service

import (
	"tat_gogogo/domain/model"
	"tat_gogogo/domain/repository"
)

type loginResult model.Result

/*
ResultService contains service of Result
*/
type ResultService struct {
	repo repository.ResultRepository
}

/*
NewResultService init a new NewResultService
@parameter: repo repository.ResultRepository
$return: *ResultService
*/
func NewResultService(repo repository.ResultRepository) *ResultService {
	return &ResultService{repo: repo}
}
