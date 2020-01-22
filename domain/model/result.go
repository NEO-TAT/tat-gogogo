package model

/*
Result is the result of Login response
success: is login successed
status: the status of response
data: the login result
*/
type Result struct {
	success bool
	status  int
	data    interface{}
}

/*
NewResult init a new result with
*/
func NewResult(success bool, status int, data interface{}) (result *Result) {
	return &Result{
		success: success,
		status: status,
		data: data,
	}
}

/*
GetSuccess is a getter of success from Result
*/
func (result Result) GetSuccess() (success bool) {
	return result.success
}

/*
GetStatus is a getter of status from Result
*/
func (result Result) GetStatus() (status int) {
	return result.status
}

/*
GetData is a getter of data from Result
*/
func (result Result) GetData() (data interface{}) {
	return result.data
}
