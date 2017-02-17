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
	if strings.Contains(post.ProfileLink.Name, "Dhamma by Ajahn Jayasaro") {
		tmplpath = "rsttemplate/jaya-en.rst"
		filename = strings.Replace(strings.ToLower(post.Title), " ", "-", -1) + "-dhamma-by-ajahn-jayasaaro%en.rst"
		return
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
