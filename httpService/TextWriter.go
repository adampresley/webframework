package httpService

import (
	"fmt"
	"net/http"
)

/*
WriteText writes text to a ResponseWriter with a content type of
text/plain
*/
func WriteText(writer http.ResponseWriter, text string, code int) {
	writer.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	writer.WriteHeader(code)
	fmt.Fprintf(writer, text)
}
