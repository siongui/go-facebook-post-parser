package parsefb

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	"testing"
)

func TestParseTime(t *testing.T) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(testhtml1))
	if err != nil {
		t.Error(err)
		return
	}

	timestamp, err := GetTimeStamp(doc)
	if err != nil {
		t.Error(err)
		return
	}

	if timestamp != "2017-02-16T07:00:00+08:00" {
		t.Error("bad timestamp: " + timestamp)
		return
	}
}
