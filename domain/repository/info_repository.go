package repository

import (
	"tat_gogogo/domain/model"

	"github.com/PuerkitoBio/goquery"
)

/*
InfoRepository declares repo of info
*/
type InfoRepository interface {
	GetInfoByRows(rows *goquery.Selection) *model.Info
}
