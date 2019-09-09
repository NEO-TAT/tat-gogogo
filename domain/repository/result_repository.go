package repository

import (
	"encoding/json"
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

type resultRepository struct{}

/*
NewResultRepository init a resultRepository
*/
func NewResultRepository() ResultRepository {
	return &resultRepository{}
}

/*
GetLoginResultByResponse handle response and get login result
*/
func (r *resultRepository) GetLoginResultByResponse(resp *http.Response) *model.Result {
	defer resp.Body.Close()

	var data map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&data)

	statusCode := 200
	isSuccess := data["success"].(bool)
	message := "登入成功"

	if !isSuccess {
		statusCode = 401
		message = "帳號或密碼錯誤，請重新輸入。"
	}

	return model.NewResult(isSuccess, statusCode, message)
}

/*
GetCurriculumResult get curriculum result by curriculums
*/
func (r *resultRepository) GetCurriculumResult(curriculums []model.Curriculum) *model.Result {
	return model.NewResult(true, 200, curriculums)
}

/*
GetCurriculumCorseResult get curriculum course result by info
*/
func (r *resultRepository) GetCurriculumCorseResult(info *model.Info) *model.Result {
	return model.NewResult(true, 200, info)
}

/*
GetNoDataResult get no data result
*/
func (r *resultRepository) GetNoDataResult() *model.Result {
	return model.NewResult(false, 400, "查無該學年或學期資料")
}
