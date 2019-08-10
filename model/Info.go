package model

/*
Info stores Curriculum"s info
*/
type Info struct {
	HasNoPeriodsCourses bool     `json:"hasNoPeriodsCourses"`
	HasSaturdayCourses  bool     `json:"hasSaturdayCourses"`
	HasSundayCourses    bool     `json:"hasSundayCourses"`
	Courses             []Course `json:"courses"`
}
