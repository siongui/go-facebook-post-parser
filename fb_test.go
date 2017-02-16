package parsefb

import (
	"testing"
)

func TestParse(t *testing.T) {
	//url := "https://www.facebook.com/jayasaro.panyaprateep.org/photos/a.318290164946343.68815.318196051622421/1119567364818615/?type=3"
	//url := "https://www.facebook.com/jayasaro.panyaprateep.org/posts/1095007907274561:0"
	//url := "https://www.facebook.com/jayasaro.panyaprateep.org/photos/a.318290164946343.68815.318196051622421/1119561951485823/?type=3"
	//result, err := Parse(url)

	post, err := ParsePost(testhtml1, "")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(post)

	post, err = ParsePost(testhtml2, "")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(post)

	post, err = ParsePost(testhtml3, "")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(post)
}
