package parsefb

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
)

func GetImageUrl(doc *goquery.Document) (string, error) {
	s := QuerySelector(doc, "img.scaledImageFitHeight")
	if s.Length() == 0 {
		s = QuerySelector(doc, "img.scaledImageFitWidth")
	}

	url, ok := s.Attr("src")
	if !ok {
		return "", errors.New("cannot find image url")
	}

	return url, nil
}
