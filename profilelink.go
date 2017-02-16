package parsefb

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
)

type ProfileLink struct {
	Name string
	Url  string
}

func GetProfileLink(doc *goquery.Document) (*ProfileLink, error) {
	s := QuerySelector(doc, "a.profileLink")
	if s.Length() == 0 {
		s = QuerySelector(doc, "span.fwb.fcg > a")
	}

	pl := ProfileLink{}

	pl.Name = s.Text()
	if pl.Name == "" {
		return nil, errors.New("cannot find name of profile link")
	}

	url, ok := s.Attr("href")
	if !ok {
		return nil, errors.New("cannot find url of profile link")
	}
	pl.Url = url

	return &pl, nil
}
