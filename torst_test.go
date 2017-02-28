package parsefb

import (
	"testing"
)

func TestToRst(t *testing.T) {
	url := "https://www.facebook.com/www.masterchingche.org/posts/1929103373990859:0"
	post, err := Parse(url)
	if err != nil {
		t.Error(err)
		return
	}

	post.Title = "要讓你念佛的善根增長廣大，你就必須提升你修學的層次!"
	post.Summary = "所以，你應該要勸你的心回到你的家，「把心帶回家」，這是第一件要做的事情。"
	post.PostUrl = `<iframe src="https://www.facebook.com/plugins/post.php?href=https%3A%2F%2Fwww.facebook.com%2Fwww.masterchingche.org%2Fposts%2F1929103373990859%3A0&width=500" width="500" height="525" style="border:none;overflow:hidden" scrolling="no" frameborder="0" allowTransparency="true"></iframe>`
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
