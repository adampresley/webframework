package webpage

import "net/http"

/*
ILayout defines an interface for layouts to adhere to
*/
type ILayout interface {
	LoadLayoutFile(fileName string) error
	LoadLayoutString(contents []byte) error
	RenderViewFile(fileName string, context interface{}) (string, error)
	RenderViewFilef(writer http.ResponseWriter, fileName string, context interface{}) error
	RenderViewString(contents []byte, context interface{}) (string, error)
	RenderViewStringf(writer http.ResponseWriter, contents []byte, context interface{}) error
}
