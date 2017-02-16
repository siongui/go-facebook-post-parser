package parsefb

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type FBPostData struct {
	PostUrl     string
	TimeStamp   string
	ProfileLink *ProfileLink
	ImageUrl    string
	Content     string
	Summary     string
	Title       string
}

func ParsePost(s, posturl string) (*FBPostData, error) {
	fb := FBPostData{PostUrl: posturl}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(s))
	if err != nil {
		return &fb, err
	}

	fb.TimeStamp, err = GetTimeStamp(doc)
	if err != nil {
		return &fb, err
	}

	fb.ProfileLink, err = GetProfileLink(doc)
	if err != nil {
		return &fb, err
	}

	fb.ImageUrl, err = GetImageUrl(doc)
	if err != nil {
		return &fb, err
	}

	fb.Content, err = GetContent(doc)
	if err != nil {
		return &fb, err
	}

	return &fb, nil
}

func Parse(url string) (*FBPostData, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	// If not login, post looks like
	// <div class="hidden_elem"><code id="u_0_f"><!-- ... --></code></div>
	s := QuerySelector(doc, "div.hidden_elem > code")
	cmt, err := s.Html()
	if err != nil {
		return nil, err
	}
	return ParsePost(cmt[5:len(cmt)-4], url)
}
