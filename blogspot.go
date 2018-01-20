package parsefb

import (
	"bufio"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"strings"
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
	t := QuerySelector(doc, "h3.post-title")
	return strings.TrimSpace(t.Text()), nil
}

func GetBlogspotContent(doc *goquery.Document) (string, error) {
	c := QuerySelector(doc, "div.post-body")

	s, err := c.Html()
	if err != nil {
		return "", err
	}

	var lines []string

	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, "  "+scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return strings.Join(lines, "\n"), nil
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
		return strings.TrimSpace(d), nil
	}

	return "", errors.New("cannot find summary")
}

func GetBlogspotAuthor(doc *goquery.Document) (string, error) {
	a := QuerySelector(doc, "span.post-author > span.fn")
	return a.Text(), nil
}

func GetBlogspotTags(doc *goquery.Document) (string, error) {
	s := doc.Find("span.post-labels > a")
	labels := ""
	s.Each(func(_ int, l *goquery.Selection) {
		if labels != "" {
			labels += ", "
		}
		labels += l.Text()
	})
	return labels, nil
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

	bs.Tags, err = GetBlogspotTags(doc)
	if err != nil {
		return &bs, err
	}

	return &bs, nil
}
