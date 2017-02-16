package parsefb

import (
	"bytes"
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
