package httpService

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

/*
WriteJSON serializes an object to JSON and writes it
to the specified writer stream with an HTTP code. The object
is of type interface{} so can technically be anything. A struct or
map would be the usual type of item to serialize.
*/
func WriteJSON(writer http.ResponseWriter, object interface{}, code int) {
	jsonBytes, _ := json.Marshal(object)
	content := strings.Replace(string(jsonBytes), "%", "%%", -1)

	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(code)
	fmt.Fprintf(writer, content)
}
