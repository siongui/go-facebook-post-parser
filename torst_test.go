package parsefb

import (
	"testing"
)

func TestToRst(t *testing.T) {
	url := "https://www.facebook.com/182989118504002/photos/a.182994541836793.46489.182989118504002/912713032198270/"
	post, err := Parse(url)
	if err != nil {
		t.Error(err)
		return
	}

	post.Title = "คำพูดล้ำค่า"
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
