package parsefb

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
)

func GetBlogspotTimeStamp(doc *goquery.Document) (string, error) {
	abbr := QuerySelector(doc, "a.timestamp-link > abbr")
	t, ok := abbr.Attr("title")
	if ok {
		return t, nil
	}

	return "", errors.New("cannot find timestamp")
}

func GetBlogspotTitle(doc *goquery.Document) (string, error) {
	t := QuerySelector(doc, "h3.post-title > a")
	return t.Text(), nil
}

func GetBlogspotContent(doc *goquery.Document) (string, error) {
	c := QuerySelector(doc, "div.post-body")
	return c.Html()
}

func GetBlogspotUrl(doc *goquery.Document) (string, error) {
	meta := QuerySelector(doc, "meta[property='og:url']")
	u, ok := meta.Attr("content")
	if ok {
		return u, nil
	}

	return "", errors.New("cannot find url")
}

func GetBlogspotSummary(doc *goquery.Document) (string, error) {
	meta := QuerySelector(doc, "meta[property='og:description']")
	d, ok := meta.Attr("content")
	if ok {
		return d, nil
	}

	return "", errors.New("cannot find summary")
}

func GetBlogspotAuthor(doc *goquery.Document) (string, error) {
	a := QuerySelector(doc, "span.post-author > span.fn")
	return a.Text(), nil
}

func ParseBlogspotPost(doc *goquery.Document) (*FBPostData, error) {
	bs := FBPostData{}
	var err error

	bs.TimeStamp, err = GetBlogspotTimeStamp(doc)
	if err != nil {
		return &bs, err
	}

	bs.Title, err = GetBlogspotTitle(doc)
	if err != nil {
		return &bs, err
	}

	bs.Content, err = GetBlogspotContent(doc)
	if err != nil {
		return &bs, err
	}

	bs.PostUrl, err = GetBlogspotUrl(doc)
	if err != nil {
		return &bs, err
	}

	bs.Summary, err = GetBlogspotSummary(doc)
	if err != nil {
		return &bs, err
	}

	bs.Author, err = GetBlogspotAuthor(doc)
	if err != nil {
		return &bs, err
	}

	return &bs, nil
}
