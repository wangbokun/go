package template

import (
	"bytes"
	"strings"
	"text/template"
)
// use:
//     data, err := util.RenderTemplate("templates/xxx.template", &res)
//    if err != nil {
// 	   return err
//    }  
func RenderTemplate(file string, v interface{}) (string, error) {
	t, err := template.ParseFiles(file)
	if err != nil {
		return "", err
	}
	t.Option("missingkey=zero")
	var buf bytes.Buffer
	err = t.Execute(&buf, v)
	return buf.String(), err
}

type Template map[string]interface{}

func (t *Template) Up(s string) string {
	return strings.ToUpper(s)
}
