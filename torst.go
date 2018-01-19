package parsefb

import (
	"bytes"
	"io"
	"os"
	"strings"
	"text/template"
)

func ToreStructuredText(post *FBPostData, tmplpath string) (string, error) {
	tmpl, err := template.ParseFiles(tmplpath)
	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	err = tmpl.Execute(&b, &post)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

func GetTemplatePath(post *FBPostData) (tmplpath, filename string) {
	if strings.Contains(post.Content, "懺公上人開示") {
		tmplpath = "rsttemplate/master-chan-yun-zh.rst"
		filename = "master-chan-yun%zh.rst"
		return
	}
	if strings.Contains(post.Content, "淨空法師親述") {
		tmplpath = "rsttemplate/master-chin-kung-zh.rst"
		filename = "master-chin-kung%zh.rst"
		return
	}
	if strings.Contains(post.Content, "良因曰") {
		tmplpath = "rsttemplate/master-liangyin%zh.rst"
		filename = "master-liangyin%zh.rst"
		return
	}
	if strings.Contains(post.Content, "聖嚴法師") {
		tmplpath = "rsttemplate/shengyen-zh.rst"
		filename = "master-sheng-yen%zh.rst"
		return
	}

	prefix := strings.Replace(strings.ToLower(post.Title), " ", "-", -1)
	if strings.Contains(post.Author, "Shen-Fu Tsai") {
		tmplpath = "rsttemplate/blogspot-sftsai-en.rst"
		filename = prefix + "%en.rst"
		return
	}
	if strings.Contains(post.ProfileLink.Name, "Dhamma by Ajahn Jayasaro") {
		if strings.Contains(post.Content, "ปิยสีโลภิกขุ") {
			tmplpath = "rsttemplate/jaya-th.rst"
			filename = prefix + "-dhamma-by-ajahn-jayasaaro%th.rst"
		} else {
			tmplpath = "rsttemplate/jaya-en.rst"
			filename = prefix + "-dhamma-by-ajahn-jayasaaro%en.rst"
		}
		return
	}
	if strings.Contains(post.ProfileLink.Url, "AjahnSuchartAbhijato") {
		tmplpath = "rsttemplate/suchart-en.rst"
		filename = "ajahn-suchart%en.rst"
		return
	}
	if strings.Contains(post.ProfileLink.Url, "masterchingche") {
		tmplpath = "rsttemplate/jingjie-zh.rst"
		filename = "master-jingjie%zh.rst"
		return
	}
	if strings.Contains(post.ProfileLink.Name, "法鼓山") {
		tmplpath = "rsttemplate/shengyen-zh.rst"
		filename = "master-sheng-yen%zh.rst"
	}
	if strings.Contains(post.Content, "ชยสาโร") {
		tmplpath = "rsttemplate/jaya2-th.rst"
		filename = "ajahn-jayasaaro%th.rst"
	}

	return
}

func SaveRst(filepath, rst string) error {
	fo, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer fo.Close()

	_, err = io.Copy(fo, strings.NewReader(rst))
	if err != nil {
		return err
	}
	return nil
}
