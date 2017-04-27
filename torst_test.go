package parsefb

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
	"strings"
	"testing"
)

type YamlData struct {
	Url      string `json:"url"`
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	PostUrl  string `json:"posturl"`
	ImageUrl string `json:"imageurl"`
}

func FacebookIframeWidthAuto(ic string) string {
	if strings.HasPrefix(ic, `<iframe src="https://www.facebook.com/plugins/post.php`) {
		return strings.Replace(strings.Replace(ic, "width=500", "width=auto", 1), `width="500"`, `width="auto"`, 1)
	}
	return ic
}

func YamlToStruct(path string) (td YamlData, err error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(b, &td)
	if err != nil {
		return
	}

	return
}

func TestToRst(t *testing.T) {
	td, err := YamlToStruct("post.yaml")
	if err != nil {
		t.Error(err)
		return
	}

	post, err := Parse(td.Url)
	if err != nil {
		t.Error(err)
		return
	}

	post.Title = td.Title
	post.Summary = td.Summary
	post.PostUrl = FacebookIframeWidthAuto(td.PostUrl)
	post.ImageUrl = td.ImageUrl
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
