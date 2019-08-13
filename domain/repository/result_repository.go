package repository

import (
	"encoding/json"
	"net/http"
	"tat_gogogo/domain/model"
)

/*
ResultRepository declare repo of result
*/
type ResultRepository interface {
	GetLoginResultByResponse(resp *http.Response) model.Result
	GetCurriculumResult(cirriculums []model.Curriculum) *model.Result
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
@parameter: *http.Response
@return: model.result
*/
func (r *resultRepository) GetLoginResultByResponse(resp *http.Response) model.Result {
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

	return *model.NewResult(isSuccess, statusCode, message)
}

func (r *resultRepository) GetCurriculumResult(curriculums []model.Curriculum) *model.Result {
	return model.NewResult(true, 201, curriculums)
}
