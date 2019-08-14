package repository

import (
	"log"
	"strings"
	"tat_gogogo/domain/model"
	"tat_gogogo/utilities/arrutil"
	"tat_gogogo/utilities/decoder"

	"github.com/PuerkitoBio/goquery"
)

var (
	columnMap = map[int]string{
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
InfoRepository declares repo of info
*/
type InfoRepository interface {
	GetInfoByRows(rows *goquery.Selection) *model.Info
}

type infoRepository struct {
	hasNoPeriodsCourses bool
	hasSaturdayCourses  bool
	hasSundayCourses    bool
	courses             []model.Course
}

/*
NewInfoRepository init a infoRepository
*/
func NewInfoRepository() InfoRepository {
	return &infoRepository{
		hasNoPeriodsCourses: false,
		hasSaturdayCourses:  false,
		hasSundayCourses:    false,
		courses:             []model.Course{},
	}
}

/*
GetInfoByRows get info by selection
@parameter: *goquery.Selection
@return: *model.Info
*/
func (i *infoRepository) GetInfoByRows(rows *goquery.Selection) *model.Info {
	rows.Each(func(rowIndex int, row *goquery.Selection) {
		indexes := []int{0, 1, 2, rows.Length() - 1}
		if arrutil.IntIndexOf(indexes, rowIndex) == -1 {
			periods := make([]string, 7)

			instructor := []string{}
			classroom := []string{}

			course := model.Course{
				Instructor: instructor,
				Periods:    periods,
				Classroom:  classroom,
				ID:         "",
				Name:       "",
			}

			columns := row.Find("td")

			parseColumns(columns, &i.hasSundayCourses, &i.hasSaturdayCourses, &course)
			organizeInfo(&i.courses, course, &i.hasNoPeriodsCourses)
		}
	})

	return model.NewInfo(
		i.courses,
		i.hasNoPeriodsCourses,
		i.hasSaturdayCourses,
		i.hasSundayCourses,
	)
}

func parseColumns(
	columns *goquery.Selection,
	hasSundayCourses *bool,
	hasSaturdayCourses *bool,
	course *model.Course,
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
	course *model.Course,
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
	course *model.Course,
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
	course *model.Course,
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

func organizeInfo(courses *[]model.Course, course model.Course, hasNoPeriodsCourses *bool) {
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
