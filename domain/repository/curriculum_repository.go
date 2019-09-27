package repository

import (
	"tat_gogogo/domain/model"

	"github.com/PuerkitoBio/goquery"
)

/*
CurriculumRepository declare repo of curriculum
*/
type CurriculumRepository interface {
	ParseCurriculums(doc *goquery.Document) []model.Curriculum
}
