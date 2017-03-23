package parsefb

import (
	"testing"
)

func TestToRst(t *testing.T) {
	url := "https://www.facebook.com/ddmbathai/posts/1470651426298851:0"
	post, err := Parse(url)
	if err != nil {
		t.Error(err)
		return
	}

	post.Title = "先接納別人才能溝通"
	post.Summary = "真正的溝通一定要先問對方有什麼困難？有什麼需求？然後再看自己能幫上什麼忙，不要一廂情願的要對方接受自己的做法。"
	post.PostUrl = `<iframe src="https://www.facebook.com/plugins/post.php?href=https%3A%2F%2Fwww.facebook.com%2Fddmbathai%2Fposts%2F1470651426298851%3A0&width=500" width="500" height="630" style="border:none;overflow:hidden" scrolling="no" frameborder="0" allowTransparency="true"></iframe>`
	post.ImageUrl = "https://scontent-tpe1-1.xx.fbcdn.net/v/t31.0-8/17492731_1470651426298851_7805806105967098325_o.jpg?oh=0ee779f50248d0609523ce1371c5abde&oe=595A6518"
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
