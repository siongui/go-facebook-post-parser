package parsefb

import (
	"testing"
)

func TestToRstJaya(t *testing.T) {
	url := "https://www.facebook.com/jayasaro.panyaprateep.org/photos/a.318290164946343.68815.318196051622421/1119561951485823/?type=3"
	post, err := Parse(url)
	if err != nil {
		t.Error(err)
		return
	}

	rst, err := ToreStructuredText(post, "rsttemplate/jaya-en.rst")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(rst)
}
