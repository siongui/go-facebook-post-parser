package parsefb

import (
	"github.com/ghodss/yaml"
	"github.com/siongui/responsive-embed-generator"
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

	if td.Title != "" {
		post.Title = td.Title
	}
	if td.Summary != "" {
		post.Summary = td.Summary
	}
	if td.PostUrl != "" {
		post.PostUrl, _ = regen.GetResponsiveFbPhotoCode(td.PostUrl)
	}
	if td.ImageUrl != "" {
		post.ImageUrl = td.ImageUrl
	}
	tmplpath, filename := GetTemplatePath(post)
	rst, err := ToreStructuredText(post, tmplpath)
	if err != nil {
		t.Error(err)
		return
	}

	rst2, err := SetRstDocTitleWidth(rst)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(rst2)

	err = SaveRst(filename, rst2)
	if err != nil {
		t.Error(err)
	}
}
