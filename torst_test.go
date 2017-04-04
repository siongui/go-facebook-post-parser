package parsefb

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
	"testing"
)

type YamlData struct {
	Url      string `json:"url"`
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	PostUrl  string `json:"posturl"`
	ImageUrl string `json:"imageurl"`
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
	post.PostUrl = td.PostUrl
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
