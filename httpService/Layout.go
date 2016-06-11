package httpService

import (
	"errors"
	"html/template"
	"net/http"
)

/*
Layout represents a directory of template files and pre-loaded bodies
*/
type Layout struct {
	BaseDirectory     string
	TemplateFileNames []string
	TemplateBodies    map[string]string
}

func NewLayout(baseDirectory string, templateFileNames []string) (Layout, error) {
	result := Layout{TemplateFileNames: templateFileNames}
	result.BaseDirectory = baseDirectory
	result.TemplateBodies = make(map[string]string)

	for _, templateName := range templateFileNames {
		html, err := LoadHTML(baseDirectory + templateName)
		if err != nil {
			return result, err
		}

		result.TemplateBodies[templateName] = string(html)
	}

	return result, nil
}

func (this Layout) parseTemplates(body string, data interface{}) (*template.Template, error) {
	t := template.New("layout")
	var err error

	for _, html := range this.TemplateBodies {
		t, err = t.Parse(html)
		if err != nil {
			return t, err
		}
	}

	t, err = t.Parse(body)
	return t, nil
}

func (this Layout) RenderView(writer http.ResponseWriter, pageName string, data interface{}) error {
	writer.Header().Set("Content-Type", "text/html; charset=UTF-8")

	html, err := LoadHTML(this.BaseDirectory + pageName)
	if err != nil {
		return err
	}

	t, err := this.parseTemplates(string(html), data)
	if err != nil {
		return err
	}

	if t == nil {
		return errors.New("Empty template: " + pageName)
	}

	return t.Execute(writer, data)
}
