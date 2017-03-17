package parsefb

import (
	"testing"
)

func TestToRst(t *testing.T) {
	url := "https://www.facebook.com/DDMCHAN/posts/1452893918100466:0"
	post, err := Parse(url)
	if err != nil {
		t.Error(err)
		return
	}

	post.Title = "禪悟與靜坐"
	post.Summary = "所謂執著的意思是什麼呢？就是當你面對任何人、物、事的時候，首先強調「我」看到了什麼，而加進價值的判斷，那就是執著。"
	post.PostUrl = `<iframe src="https://www.facebook.com/plugins/post.php?href=https%3A%2F%2Fwww.facebook.com%2FDDMCHAN%2Fposts%2F1452893918100466%3A0&width=500" width="500" height="500" style="border:none;overflow:hidden" scrolling="no" frameborder="0" allowTransparency="true"></iframe>`
	post.ImageUrl = "https://scontent-tpe1-1.xx.fbcdn.net/v/t31.0-8/17311354_1452893918100466_7254377630723103173_o.jpg?oh=056572cd3360c05448834fbe2b5a406d&oe=5924F91F"
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
