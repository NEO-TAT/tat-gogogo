package curriculum

import (
	"log"
	"net/http"
	"net/url"
	"strings"
	"tat_gogogo/configs"
	"tat_gogogo/crawler/portal"

	"tat_gogogo/utilities/decoder"

	"github.com/PuerkitoBio/goquery"
)

/*
Curriculum stores Year and semester of curriculum
*/
type Curriculum struct {
	Year     string `json:"year"`
	Semester string `json:"semester"`
}

/*
Info stores Curriculum"s info
*/
type Info struct {
	HasNoPeriodsCourses bool     `json:"hasNoPeriodsCourses"`
	HasSaturdayCourses  bool     `json:"hasSaturdayCourses"`
	HasSundayCourses    bool     `json:"hasSundayCourses"`
	Courses             []Course `json:"courses"`
}

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

var (
	config, configError = configs.New()
	client              = portal.NewClient()
	columnMap           = map[int]string{
		0:  "id",
		1:  "name",
		6:  "instructor",
		8:  "periodsOfSunday",
		9:  "periodsOfMonday",
		10: "periodsOfTuesday",
		11: "periodsOfWednesday",
		12: "periodsOfThursday",
		13: "periodsOfFriday",
		14: "periodsOfSaturday",
		15: "classroom",
	}
)

/*
GetCurriculums return the curriculums
the default targetStudentID will be self
*/
func GetCurriculums(
	studentID string,
	password string,
	targetStudentID string,
) (curriculumResult portal.Result, err error) {
	if configError != nil {
		log.Panicln(configError)
		return portal.Result{}, configError
	}

	curriculumLoginResult, err := loginCurriculum(studentID, password)
	if err != nil {
		log.Panicln(err)
		return portal.Result{}, err
	}

	if !curriculumLoginResult.Success {
		return curriculumLoginResult, nil
	}

	if targetStudentID == "" {
		return handleCurriculumRequest(studentID)
	}

	return handleCurriculumRequest(targetStudentID)
}

func handleCurriculumRequest(targetStudentID string) (result portal.Result, err error) {
	form := url.Values{
		"code":   {targetStudentID},
		"format": {"-3"},
	}

	curriculumRequest, err := http.NewRequest("POST", config.CoureseSystem.Select, strings.NewReader(form.Encode()))
	if err != nil {
		log.Panicln(err)
		return portal.Result{}, err
	}
	curriculumRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	curriculumsResp, err := client.Do(curriculumRequest)
	if err != nil {
		log.Panicln(err)
		return portal.Result{}, err
	}

	curriculumDoc, err := goquery.NewDocumentFromReader(curriculumsResp.Body)
	if err != nil {
		log.Panicln(err)
		return portal.Result{}, err
	}

	return parseCurriculums(curriculumDoc), nil
}

func parseCurriculums(doc *goquery.Document) (result portal.Result) {
	var curriculums []Curriculum

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		if href, ok := s.Attr("href"); ok {
			splits := strings.Split(href, "&")
			year := strings.Replace(splits[2], "year=", "", 1)
			sem := strings.Replace(splits[3], "sem=", "", 1)

			curriculum := Curriculum{Year: year, Semester: sem}
			curriculums = append(curriculums, curriculum)
		}
	})

	return portal.Result{Data: curriculums, Status: 200, Success: true}
}

func loginCurriculum(
	studentID string,
	password string,
) (loginCourseResult portal.Result, err error) {
	loginResult, err := portal.Login(client, studentID, password)

	if err != nil {
		log.Panicln(err)
		return portal.Result{}, err
	}

	if loginResult.Status != 200 {
		return loginResult, nil
	}

	req, err := http.NewRequest("POST", config.Portal.SsoLoginCourseSystem, nil)
	if err != nil {
		log.Panicln(err)
		return portal.Result{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
		return portal.Result{}, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Panicln(err)
		return portal.Result{}, err
	}

	return accessCourse(doc)
}

func accessCourse(doc *goquery.Document) (loginCourseResult portal.Result, err error) {
	form := parseFormBy(doc)

	bufferRequest, err := http.NewRequest("POST", config.CoureseSystem.MainPage, strings.NewReader(form.Encode()))
	if err != nil {
		log.Panicln(err)
		return portal.Result{Success: false, Status: 500}, err
	}
	bufferRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	buffer, err := client.Do(bufferRequest)
	if err != nil {
		log.Panicln(err)
		return portal.Result{Success: false, Status: 500}, err
	}

	return parseLoginCurriculumResult(buffer)
}

func parseFormBy(doc *goquery.Document) (from url.Values) {
	form := url.Values{}
	doc.Find("input[type=hidden]").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("name")
		value, _ := s.Attr("value")
		form.Add(name, value)
	})

	return form
}

func parseLoginCurriculumResult(
	buffer *http.Response,
) (loginCourseResult portal.Result, err error) {
	defer buffer.Body.Close()

	courseDoc, err := goquery.NewDocumentFromReader(buffer.Body)
	if err != nil {
		log.Panicln(err)
		return portal.Result{Success: false, Status: 500}, err
	}

	rawLast := courseDoc.Find("body a").Last().Text()
	last, err := decoder.DecodeToBig5(rawLast)
	if err != nil {
		log.Println(err)
		return portal.Result{Success: false, Status: 500}, err
	}

	var status int
	var message string

	isLoginSuccess := last == "依 [學號]／[課號] 查詢選課表"

	if !isLoginSuccess {
		status = 401
		message = "登入課程系統失敗"
	} else {
		status = 200
		message = "登入課程系統成功"
	}

	return portal.Result{Success: isLoginSuccess, Status: status, Data: message}, nil
}
