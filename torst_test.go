package parsefb

import (
	"testing"
)

func TestToRst(t *testing.T) {
	url := "https://www.facebook.com/DDMCHAN/photos/a.169338186456052.44118.143595409030330/1392471394142719/"
	post, err := Parse(url)
	if err != nil {
		t.Error(err)
		return
	}

	post.Title = "妄念不起．萬緣不拒"
	tmplpath, filename := GetTemplatePath(post)
	rst, err := ToreStructuredText(post, tmplpath)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(rst)

	err = SaveRst(filename, rst)
	if err != nil {
		t.Error(err)
	}
}
