package parsefb

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"time"
)

func ParseTimeStamp(utime string) (string, error) {
	i, err := strconv.ParseInt(utime, 10, 64)
	if err != nil {
		return "", err
	}
	t := time.Unix(i, 0)
	return t.Format(time.RFC3339), nil
}

func GetTimeStamp(doc *goquery.Document) (string, error) {
	//s := QuerySelector(doc, "._5ptz.timestamp.livetimestamp")
	s := QuerySelector(doc, "abbr._5ptz")
	utime, ok := s.Attr("data-utime")
	if ok {
		return ParseTimeStamp(utime)
	}

	return "", errors.New("cannot find timestamp")
}
