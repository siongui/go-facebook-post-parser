package parsefb

import (
	"testing"
)

func TestToRst(t *testing.T) {
	url := "https://www.facebook.com/DDMCHAN/posts/1439508549439003:0"
	post, err := Parse(url)
	if err != nil {
		t.Error(err)
		return
	}

	post.Title = "智者無為"
	post.Summary = "智者無為，愚人自縛。"
	post.PostUrl = `<iframe src="https://www.facebook.com/plugins/post.php?href=https%3A%2F%2Fwww.facebook.com%2FDDMCHAN%2Fposts%2F1439508549439003%3A0&width=500" width="500" height="428" style="border:none;overflow:hidden" scrolling="no" frameborder="0" allowTransparency="true"></iframe>`
	post.ImageUrl = "https://scontent-tpe1-1.xx.fbcdn.net/v/t31.0-8/16904919_1439508549439003_3217797208732337712_o.jpg?oh=816801c71fc7d0c1013d2538eff7d1f4&oe=5933402D"
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
