package webpage

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/adampresley/webframework/httpService"
	"github.com/pkg/errors"
)

/*
GoLayout is a layout using Go templates
*/
type GoLayout struct {
	layout *template.Template
}

func (l *GoLayout) addHelperFunctions() {
	var funcMap = template.FuncMap{
		"eq": func(a, b interface{}) bool {
			return a == b
		},
	}

	l.layout.Funcs(funcMap)
}

/*
LoadLayoutFile loads a layout from a file
*/
func (l *GoLayout) LoadLayoutFile(fileName string) error {
	var err error
	var layoutContents []byte

	if layoutContents, err = ioutil.ReadFile(fileName); err != nil {
		return err
	}

	l.layout = template.Must(template.New("layout").Parse(string(layoutContents[:len(layoutContents)])))
	l.addHelperFunctions()
	return nil
}

/*
LoadLayoutString loads a layout from a passed-in byte array
*/
func (l *GoLayout) LoadLayoutString(contents []byte) error {
	l.layout = template.Must(template.New("layout").Parse(string(contents[:len(contents)])))
	l.addHelperFunctions()
	return nil
}

/*
RenderViewFile renders a view from a file into this layout
*/
func (l *GoLayout) RenderViewFile(fileName string, context interface{}) (string, error) {
	var viewContents []byte
	var err error
	var viewTemplate *template.Template
	stringWriter := bytes.NewBufferString("")

	if viewContents, err = ioutil.ReadFile(fileName); err != nil {
		return "", errors.Wrapf(err, "Unable to read the view file %s", fileName)
	}

	if viewTemplate, err = l.layout.Parse(string(viewContents[:len(viewContents)])); err != nil {
		return "", errors.Wrapf(err, "Unable to parse the view file %s", fileName)
	}

	if err = viewTemplate.Execute(stringWriter, context); err != nil {
		return "", errors.Wrapf(err, "Unable to render the file %s", fileName)
	}

	return stringWriter.String(), nil
}

/*
RenderViewFilef renders a view from a file into this layout, then
writes it out to the provider writer. Useful for HTTP responses
*/
func (l *GoLayout) RenderViewFilef(writer http.ResponseWriter, fileName string, context interface{}) error {
	var renderedContents string
	var err error

	if renderedContents, err = l.RenderViewFile(fileName, context); err != nil {
		return errors.Wrapf(err, "Unable to render the view file %s", fileName)
	}

	httpService.WriteHTML(writer, renderedContents, 200)
	return nil
}

/*
RenderViewString renders a view from byte array content into this layout
*/
func (l *GoLayout) RenderViewString(contents []byte, context interface{}) (string, error) {
	var err error

	viewTemplate := template.Must(template.New("view").Parse(string(contents[:len(contents)])))
	stringWriter := bytes.NewBufferString("")

	if err = viewTemplate.ExecuteTemplate(stringWriter, "layout", context); err != nil {
		return "", errors.Wrapf(err, "Unable to render the view")
	}

	return stringWriter.String(), nil
}

/*
RenderViewStringf renders a view from byte array content into this layout, and
writes it out to the provided writer. Useful for HTTP responses
*/
func (l *GoLayout) RenderViewStringf(writer http.ResponseWriter, contents []byte, context interface{}) error {
	var renderedContents string
	var err error

	if renderedContents, err = l.RenderViewString(contents, context); err != nil {
		return errors.Wrap(err, "Unable to render the view")
	}

	httpService.WriteHTML(writer, renderedContents, 200)
	return nil
}
