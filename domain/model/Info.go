package model

/*
Info stores Curriculum"s info
*/
type Info struct {
	Courses             []Course `json:"courses"`
	HasNoPeriodsCourses bool     `json:"hasNoPeriodsCourses"`
	HasSaturdayCourses  bool     `json:"hasSaturdayCourses"`
	HasSundayCourses    bool     `json:"hasSundayCourses"`
}

/*
NewInfo init a new Info
*/
func NewInfo(courses []Course, hasNoPeriodsCourses, hasSaturdayCourses, hasSundayCourses bool) (info *Info) {
	return &Info{
		Courses:             courses,
		HasNoPeriodsCourses: hasNoPeriodsCourses,
		HasSaturdayCourses:  hasSaturdayCourses,
		HasSundayCourses:    hasSundayCourses,
	}
}
