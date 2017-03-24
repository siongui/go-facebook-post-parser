package parsefb

import (
	"testing"
)

func TestToRst(t *testing.T) {
	url := "https://www.facebook.com/DDMCHAN/posts/1457133827676475:0"
	post, err := Parse(url)
	if err != nil {
		t.Error(err)
		return
	}

	post.Title = "虛與實"
	post.Summary = "我們凡夫是顛倒的，以虛為實，以實為虛，所以產生煩惱。看到好看的就貪，遇到不適意的就瞋。如果能離開虛幻的現象，便像《圓覺經》說的「知幻即離」、「離幻即覺」。就能夠得解脫、得自在。"
	post.PostUrl = `<iframe src="https://www.facebook.com/plugins/post.php?href=https%3A%2F%2Fwww.facebook.com%2FDDMCHAN%2Fposts%2F1457133827676475%3A0&width=500" width="500" height="518" style="border:none;overflow:hidden" scrolling="no" frameborder="0" allowTransparency="true"></iframe>`
	post.ImageUrl = "https://scontent-tpe1-1.xx.fbcdn.net/v/t31.0-8/17358970_1457133827676475_4569122256418795348_o.jpg?oh=5f5576f49f55a30fd910964484289e11&oe=5963C87D"
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
