package parsefb

import (
	"github.com/PuerkitoBio/goquery"
)

func GetContent(doc *goquery.Document) (string, error) {
	s := QuerySelector(doc, "div._5pbx.userContent")
	if s.Length() == 0 {
		return "no content", nil
	}

	content, err := s.Html()
	if err != nil {
		return "", err
	}

	return content, nil
}
