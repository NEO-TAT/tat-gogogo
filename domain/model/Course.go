package model

/*
Course stores the course information
*/
type Course struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Instructor []string `json:"instructor"`
	Periods    []string `json:"periods"`
	Classroom  []string `json:"classroom"`
}