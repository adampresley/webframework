package httpService

import (
	"fmt"
	"net/http"
)

/*
WriteHTML writes HTML out to a ResponseWriter
*/
func WriteHTML(writer http.ResponseWriter, html string, code int) {
	writer.Header().Set("Content-Type", "text/html; charset=UTF-8")
	writer.WriteHeader(code)
	fmt.Fprint(writer, html)
}
