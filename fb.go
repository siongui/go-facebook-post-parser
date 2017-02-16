package parsefb

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func ParsePost(s, posturl string) (string, error) {
	println(s)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(s))
	if err != nil {
		return "", err
	}

	result := ""
	result += posturl
	result += "\n"

	timestamp, err := GetTimeStamp(doc)
	if err != nil {
		return "", err
	}
	result += timestamp
	result += "\n"

	pl, err := GetProfileLink(doc)
	if err != nil {
		return "", err
	}
	result += pl.Name
	result += "\n"
	result += pl.Url
	result += "\n"

	imgurl, err := GetImageUrl(doc)
	if err != nil {
		return "", err
	}
	result += imgurl
	result += "\n"

	content, err := GetContent(doc)
	if err != nil {
		return "", err
	}
	result += content
	result += "\n"

	return result, nil
}

func Parse(url string) (string, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return "", err
	}

	// If not login, post looks like
	// <div class="hidden_elem"><code id="u_0_f"><!-- ... --></code></div>
	s := QuerySelector(doc, "div.hidden_elem > code")
	cmt, err := s.Html()
	if err != nil {
		return "", err
	}
	return ParsePost(cmt[5:len(cmt)-4], url)
}
