package httpService

import "net/http"

/*
A BasicJSONResponse allows you to return a message and a success
indicator
*/
type BasicJSONResponse struct {
	IsSuccess bool   `json:"success"`
	Message   string `json:"message"`
}

/*
NewBasicJSONResponse creates a new structure.
*/
func NewBasicJSONResponse(success bool, message string) BasicJSONResponse {
	return BasicJSONResponse{
		IsSuccess: success,
		Message:   message,
	}
}

func ErrorJSONResponse(message string) BasicJSONResponse {
	return BasicJSONResponse{
		IsSuccess: false,
		Message:   message,
	}
}

/*
BadRequest sends this basic JSON response as a 400
*/
func (response BasicJSONResponse) BadRequest(writer http.ResponseWriter) {
	BadRequest(writer, response)
}

/*
Error sends this basic JSON response as a 500
*/
func (response BasicJSONResponse) Error(writer http.ResponseWriter) {
	Error(writer, response)
}

/*
Forbidden sends this basic JSON response as a 403
*/
func (response BasicJSONResponse) Forbidden(writer http.ResponseWriter) {
	Forbidden(writer, response)
}

/*
NotFound sends this basic JSON response as a 404
*/
func (response BasicJSONResponse) NotFound(writer http.ResponseWriter) {
	NotFound(writer, response)
}

/*
Success sends this basic JSON response as a 200
*/
func (response BasicJSONResponse) Success(writer http.ResponseWriter) {
	Success(writer, response)
}
