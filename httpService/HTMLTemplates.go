package httpService

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

/*
LoadHTML loads a file with the specified name and returns a byte array of
its contents. This will automatically append the HTML extension to the
end of the page name
*/
func LoadHTML(pageName string) ([]byte, error) {
	fileName := pageName + ".html"
	return ioutil.ReadFile(fileName)
}

/*
RenderHTML takes the name of a page, loads it, and writes it to the provided
ResponseWriter.
*/
func RenderHTML(writer http.ResponseWriter, pageName string) error {
	body, err := LoadHTML(pageName)
	if err != nil {
		return err
	}

	writer.Header().Set("Content-Type", "text/html; charset=UTF-8")

	fmt.Fprintf(writer, string(body))
	return nil
}

/*
RenderTemplate takes a page name, loads it, executes it as a template against the provided
data, then writes it to the provided ResponseWriter
*/
func RenderTemplate(writer http.ResponseWriter, pageName string, data interface{}) error {
	t, err := template.ParseFiles(pageName + ".html")
	if err != nil {
		return err
	}

	err = t.Execute(writer, data)
	return err
}
