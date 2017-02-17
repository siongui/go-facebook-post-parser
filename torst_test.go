package parsefb

import (
	"testing"
)

func TestToRst(t *testing.T) {
	url := "https://www.facebook.com/jayasaro.panyaprateep.org/photos/a.318290164946343.68815.318196051622421/1120112961430722/"
	post, err := Parse(url)
	if err != nil {
		t.Error(err)
		return
	}

	post.Title = "The Buddha taught that wise people"
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
