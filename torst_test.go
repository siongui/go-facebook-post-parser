package parsefb

import (
	"testing"
)

func TestToRst(t *testing.T) {
	url := "https://www.facebook.com/DDMCHAN/photos/a.169338186456052.44118.143595409030330/1439489769440881/?type=3"
	post, err := Parse(url)
	if err != nil {
		t.Error(err)
		return
	}

	post.Title = "至道無難"
	post.Summary = "至道無難，唯嫌揀擇。但莫憎愛，洞然明白。"
	post.PostUrl = `<iframe src="https://www.facebook.com/plugins/post.php?href=https%3A%2F%2Fwww.facebook.com%2FDDMCHAN%2Fposts%2F1439489769440881%3A0&width=500" width="500" height="481" style="border:none;overflow:hidden" scrolling="no" frameborder="0" allowTransparency="true"></iframe>`
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
