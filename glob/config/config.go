package config

/*
Constants is a struct which store the const within the projsct
*/
type Constants struct {
	Host          string
	Port          string
	Portal        Portal
	CoureseSystem CoureseSystem
}

/*
Portal contains all the urls of application
*/
type Portal struct {
	Base                 string
	IndexPage            string
	Login                string
	MainPage             string
	AptreeListPage       string
	AptreeAAListPage     string
	SsoLoginCourseSystem string
}

/*
CoureseSystem contains all of the elements from course system
*/
type CoureseSystem struct {
	Base     string
	MainPage string
	Select   string
}
