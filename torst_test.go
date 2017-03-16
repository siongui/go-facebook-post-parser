package parsefb

import (
	"testing"
)

func TestToRst(t *testing.T) {
	url := "https://www.facebook.com/AjahnSuchartAbhijato/posts/746180318880060:0"
	post, err := Parse(url)
	if err != nil {
		t.Error(err)
		return
	}

	post.Title = "Just do it. Just like eating"
	post.Summary = "“Just do it. Just like eating — you have to eat every day. But some days you may not feel like eating, but you know that the result of not eating is more painful than eating. It is the same with the practice. The result of not practising is more painful than practising.”"
	post.PostUrl = `<iframe src="https://www.facebook.com/plugins/post.php?href=https%3A%2F%2Fwww.facebook.com%2FAjahnSuchartAbhijato%2Fposts%2F746180318880060%3A0&width=500" width="500" height="567" style="border:none;overflow:hidden" scrolling="no" frameborder="0" allowTransparency="true"></iframe>`
	post.ImageUrl = "https://scontent-tpe1-1.xx.fbcdn.net/v/t1.0-9/17342924_746180318880060_5263953981412877604_n.jpg?oh=ea729285092746106c70be67acc7511d&oe=5959B6DF"
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
