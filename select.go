package parsefb

import (
	"github.com/PuerkitoBio/goquery"
)

type object interface {
	Find(string) *goquery.Selection
}

func QuerySelector(s object, selector string) *goquery.Selection {
	return s.Find(selector).First()
}
