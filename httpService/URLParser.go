package httpService

import (
	"net/http"
	"strconv"
	"strings"
)

/*
SplitListOfIntsInURL takes a delimited list value in a URL key and
splits it into a slice of integers
*/
func SplitListOfIntsInURL(request *http.Request, key string, delimiter string) []int {
	sourceValue := request.URL.Query().Get(key)
	splitValues := strings.Split(sourceValue, delimiter)
	result := make([]int, 0, len(splitValues))

	var err error
	var convertedValue int

	for _, value := range splitValues {
		if convertedValue, err = strconv.Atoi(value); err != nil {
			result = append(result, 0)
		} else {
			result = append(result, convertedValue)
		}
	}

	return result
}
