package parsefb

import (
	"testing"
)

func TestToRst(t *testing.T) {
	url := "https://www.facebook.com/DDMCHAN/photos/a.169338186456052.44118.143595409030330/1392460990810426/?type=3"
	post, err := Parse(url)
	if err != nil {
		t.Error(err)
		return
	}

	post.Title = "自我的消融"
	post.Summary = "消融了自我的執著、自我的煩惱，便能顯現出無我的大智慧以及平等的大慈悲；實際上就是自我的無限自在與無限包容。"
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
