package curriculum

import (
	"log"
	"net/http"
	"net/url"
	"strings"
	"tat_gogogo/configs"
	"tat_gogogo/crawler/portal"

	"tat_gogogo/utilities/arrhelp"
	"tat_gogogo/utilities/decoder"

	"github.com/PuerkitoBio/goquery"
)

/*
Curriculum stores Year and semester of curriculum
*/
type Curriculum struct {
	Year     string `json:"year"`
	Semester string `json:"Semester"`
}

/*
Info stores Curriculum"s info
*/
type Info struct {
	HasNoPeriodsCourses bool     `json:"hasNoPeriodsCourses"`
	HasSaturdayCourses  bool     `json:"HasSaturdayCourses"`
	HasSundayCourses    bool     `json:"HasSundayCourses"`
	Courses             []Course `json:"Courses"`
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
)

/*
GetCurriculums return the curriculums
the default targetStudentID will be self
*/
func GetCurriculums(studentID string, password string, targetStudentID string) (curriculumResult portal.Result, err error) {
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

/*
GetCurriculumCourse return the course info
*/
func GetCurriculumCourse(studentID string, password string, targetStudentID string, year string, sem string) (curriculumCourseResult portal.Result, err error) {
	curriculumResult, err := GetCurriculums(studentID, password, targetStudentID)
	if err != nil {
		log.Panicln(err)
		return portal.Result{}, err
	}

	if !curriculumResult.Success {
		return curriculumResult, nil
	}

	isSameYearAndSem := false
	for _, curriculum := range curriculumResult.Data.([]Curriculum) {
		if curriculum.Year == year && curriculum.Semester == sem {
			isSameYearAndSem = true
			break
		}
	}

	if !isSameYearAndSem {
		return portal.Result{Data: "查無該學年或學期資料", Status: 400, Success: false}, nil
	}

	info, err := getCurriculumInfo(studentID, year, sem)
	if err != nil {
		log.Panicln(err)
		return portal.Result{}, err
	}

	return portal.Result{Data: info, Status: 200, Success: true}, nil
}

func handleCurriculumRequest(targetStudentID string) (result portal.Result, err error) {
	form := url.Values{
		"code":   {targetStudentID},
		"format": {"-3"},
	}

	curriculumRequest, err := http.NewRequest("POST", config.COURESESYSTEM.Select, strings.NewReader(form.Encode()))
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

func loginCurriculum(studentID string, password string) (loginCourseResult portal.Result, err error) {
	loginResult, err := portal.Login(client, studentID, password)

	if err != nil {
		log.Panicln(err)
		return portal.Result{}, err
	}

	if loginResult.Status != 200 {
		return loginResult, nil
	}

	req, err := http.NewRequest("POST", config.PORTAL.SsoLoginCourseSystem, nil)
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

	bufferRequest, err := http.NewRequest("POST", config.COURESESYSTEM.MainPage, strings.NewReader(form.Encode()))
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

func parseLoginCurriculumResult(buffer *http.Response) (loginCourseResult portal.Result, err error) {
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

func getCurriculumInfo(studentID string, year string, sem string) (info Info, err error) {
	buffer, err := handleCourseSelectRequest(studentID, year, sem)
	if err != nil {
		return Info{}, err
	}

	defer buffer.Body.Close()

	courseDoc, err := goquery.NewDocumentFromReader(buffer.Body)
	if err != nil {
		log.Panicln(err)
	}

	var courses []Course
	columnMap := map[int]string{
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

	hasNoPeriodsCourses := false
	hasSaturdayCourses := false
	hasSundayCourses := false

	rows := courseDoc.Find("table").Last().Find("tr")

	rows.Each(func(rowIndex int, row *goquery.Selection) {
		indexes := []int{0, 1, 2, rows.Length() - 1}
		if arrhelp.IntIndexOf(indexes, rowIndex) == -1 {
			instructor := []string{}
			periods := make([]string, 7)
			classroom := []string{}

			course := Course{Instructor: instructor, Periods: periods, Classroom: classroom, ID: "", Name: ""}

			columns := row.Find("td")

			columns.Each(func(columnIndex int, column *goquery.Selection) {

				if _, ok := columnMap[columnIndex]; ok {
					var element *goquery.Selection

					if column.Find("a").Length() == 0 {
						element = column
					} else {
						element = column.Find("a")
					}

					if columnIndex >= 8 && columnIndex <= 14 {
						day := columnIndex - 8

						big5Element, _ := decoder.DecodeToBig5(element.Text())
						course.Periods[day] = strings.TrimSpace(big5Element)

						if !hasSaturdayCourses && day == 6 && len(course.Periods[day]) != 0 {
							hasSaturdayCourses = true
						}

						if !hasSundayCourses && day == 0 && len(course.Periods[day]) != 0 {
							hasSundayCourses = true
						}
					} else if columnIndex == 6 || columnIndex == 15 {
						element.Each(func(i int, el *goquery.Selection) {
							switch columnMap[columnIndex] {
							case "instructor":
								big5Instructor, err := decoder.DecodeToBig5(el.Text())
								if err != nil {
									log.Panicln(err)
									break
								}
								course.Instructor = append(course.Instructor, strings.TrimSpace(big5Instructor))
							case "classroom":
								big5Classroom, err := decoder.DecodeToBig5(el.Text())
								if err != nil {
									log.Panicln(err)
									break
								}
								course.Classroom = append(course.Classroom, strings.TrimSpace(big5Classroom))
							default:
								log.Println("beyond the map", columnMap[columnIndex])
							}
						})
					} else {
						switch columnMap[columnIndex] {
						case "id":
							big5Id, err := decoder.DecodeToBig5(element.Text())
							if err != nil {
								log.Panicln(err)
								break
							}
							course.ID = strings.TrimSpace(big5Id)
						case "name":
							big5Name, err := decoder.DecodeToBig5(element.Text())
							if err != nil {
								log.Panicln(err)
								break
							}
							course.Name = strings.TrimSpace(big5Name)
						default:
							log.Println("beyond the map", columnMap[columnIndex])
						}
					}
				}
			})

			courses = append(courses, course)
			if !hasNoPeriodsCourses {
				temp := true
				for _, p := range course.Periods {
					if len(p) != 0 {
						temp = false
						break
					}
				}
				hasNoPeriodsCourses = temp
			}
		}
	})

	return Info{
			HasNoPeriodsCourses: hasNoPeriodsCourses,
			HasSaturdayCourses:  hasSaturdayCourses,
			HasSundayCourses:    hasSundayCourses,
			Courses:             courses,
		},
		nil
}

func handleCourseSelectRequest(studentID string, year string, sem string) (buffer *http.Response, err error) {
	bufferReq, err := http.NewRequest("GET", config.COURESESYSTEM.Select, nil)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	bufferReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	q := bufferReq.URL.Query()
	q.Add("format", "-2")
	q.Add("code", studentID)
	q.Add("year", year)
	q.Add("sem", sem)

	bufferReq.URL.RawQuery = q.Encode()

	buffer, err = client.Do(bufferReq)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	return buffer, nil
}
