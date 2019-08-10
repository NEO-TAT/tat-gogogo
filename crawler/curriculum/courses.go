package curriculum

import (
	"log"
	"net/http"
	"strings"
	"tat_gogogo/crawler/portal"
	"tat_gogogo/utilities/arrutil"
	"tat_gogogo/utilities/decoder"

	"github.com/PuerkitoBio/goquery"
)

/*
GetCourses return the courses from target student info
the default studentID will be self
*/
func GetCourses(
	studentID string,
	password string,
	targetStudentID string,
	year string,
	sem string,
) (curriculumCourseResult portal.Result, err error) {
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

func getCurriculumInfo(
	studentID string,
	year string,
	sem string,
) (info Info, err error) {
	buffer, err := handleCourseSelectRequest(studentID, year, sem)
	if err != nil {
		return Info{}, err
	}

	defer buffer.Body.Close()

	courseDoc, err := goquery.NewDocumentFromReader(buffer.Body)
	if err != nil {
		log.Panicln(err)
	}

	rows := courseDoc.Find("table").Last().Find("tr")

	return parseRows(rows), nil
}

func parseRows(rows *goquery.Selection) (info Info) {
	hasNoPeriodsCourses := false
	hasSaturdayCourses := false
	hasSundayCourses := false
	courses := []Course{}
	rows.Each(func(rowIndex int, row *goquery.Selection) {
		indexes := []int{0, 1, 2, rows.Length() - 1}
		if arrutil.IntIndexOf(indexes, rowIndex) == -1 {
			periods := make([]string, 7)

			instructor := []string{}
			classroom := []string{}

			course := Course{
				Instructor: instructor,
				Periods:    periods,
				Classroom:  classroom,
				ID:         "",
				Name:       "",
			}

			columns := row.Find("td")

			parseColumns(columns, &hasSundayCourses, &hasSaturdayCourses, &course)
			organizeInfo(&courses, course, &hasNoPeriodsCourses)
		}
	})

	return Info{
		HasNoPeriodsCourses: hasNoPeriodsCourses,
		HasSaturdayCourses:  hasSaturdayCourses,
		HasSundayCourses:    hasSundayCourses,
		Courses:             courses,
	}
}

func organizeInfo(courses *[]Course, course Course, hasNoPeriodsCourses *bool) {
	*courses = append(*courses, course)
	if !(*hasNoPeriodsCourses) {
		temp := true
		for _, p := range course.Periods {
			if len(p) != 0 {
				temp = false
				break
			}
		}
		*hasNoPeriodsCourses = temp
	}
}

func parseColumns(
	columns *goquery.Selection,
	hasSundayCourses *bool,
	hasSaturdayCourses *bool,
	course *Course,
) {
	columns.Each(func(columnIndex int, column *goquery.Selection) {
		if _, ok := columnMap[columnIndex]; ok {
			var element *goquery.Selection

			if column.Find("a").Length() == 0 {
				element = column
			} else {
				element = column.Find("a")
			}

			if columnIndex >= 8 && columnIndex <= 14 {
				handlePeriods(columnIndex, element, hasSundayCourses, hasSaturdayCourses, course)
			} else if columnIndex == 6 || columnIndex == 15 {
				handleInstructorAndClassroom(element, columnIndex, course)
			} else {
				handleIDAndName(element, columnIndex, course)
			}
		}
	})
}

func handlePeriods(
	columnIndex int,
	element *goquery.Selection,
	hasSundayCourses *bool,
	hasSaturdayCourses *bool,
	course *Course,
) {
	day := columnIndex - 8
	big5Element, _ := decoder.DecodeToBig5(element.Text())
	course.Periods[day] = strings.TrimSpace(big5Element)

	if !(*hasSaturdayCourses) && day == 6 && len(course.Periods[day]) != 0 {
		*hasSaturdayCourses = true
	}

	if !(*hasSundayCourses) && day == 0 && len(course.Periods[day]) != 0 {
		*hasSundayCourses = true
	}
}

func handleInstructorAndClassroom(
	element *goquery.Selection,
	columnIndex int,
	course *Course,
) {
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
}

func handleIDAndName(
	element *goquery.Selection,
	columnIndex int,
	course *Course,
) {
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

func handleCourseSelectRequest(
	studentID string,
	year string,
	sem string,
) (buffer *http.Response, err error) {
	bufferReq, err := http.NewRequest("GET", config.CoureseSystem.Select, nil)
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
