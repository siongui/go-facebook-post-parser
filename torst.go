package parsefb

import (
	"bytes"
	"fmt"
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
	prefix := strings.Replace(strings.ToLower(post.Title), " ", "-", -1)
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

	_, err = fmt.Fprintf(fo, rst)
	if err != nil {
		return err
	}
	return nil
}
