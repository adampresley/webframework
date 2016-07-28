package httpService

import (
	"net/http"
)

/*
AccessDenied is a convenience method. This will set the HTTP status
code to 403 and send back a BasicJSONResponse with a message
of Access Denied
*/
func AccessDenied(writer http.ResponseWriter) {
	NewBasicJSONResponse(false, "Access Denied").Forbidden(writer)
}

/*
BadRequest is a convenience method. This will set the HTTP status
code to 400. This is useful for telling a caller that the
request is malformed.
*/
func BadRequest(writer http.ResponseWriter, response interface{}) {
	WriteJSON(writer, response, 400)
}

/*
Error is a convenience method. This will set the HTTP status
code to 500. This is useful for indicating some type of
error ocurred during processing.
*/
func Error(writer http.ResponseWriter, response interface{}) {
	WriteJSON(writer, response, 500)
}

/*
Forbidden is a convenience method. This will set the HTTP status
code to 403. This is useful in telling a caller that some
type of authentication failed.
*/
func Forbidden(writer http.ResponseWriter, response interface{}) {
	WriteJSON(writer, response, 403)
}

/*
JSONBadRequest is a shortcut to send a bad request with a basic JSON object
*/
func JSONBadRequest(writer http.ResponseWriter, message string) {
	NewBasicJSONResponse(false, message).BadRequest(writer)
}

/*
JSONError is a shortcut to send an error with a basic JSON object
*/
func JSONError(writer http.ResponseWriter, message string) {
	NewBasicJSONResponse(false, message).Error(writer)
}

/*
JSONForbidden is a shortcut to send a forbidden response with a basic JSON object
*/
func JSONForbidden(writer http.ResponseWriter, message string) {
	NewBasicJSONResponse(false, message).Forbidden(writer)
}

/*
JSONNotFound is a shortcut to send a not found message with a basic JSON object
*/
func JSONNotFound(writer http.ResponseWriter, message string) {
	NewBasicJSONResponse(false, message).NotFound(writer)
}

/*
JSONSuccess is a shortcut to send a 200 with a basic JSON object
*/
func JSONSuccess(writer http.ResponseWriter, message string) {
	NewBasicJSONResponse(true, message).Success(writer)
}

/*
NotFound is a convenience method. This will set the HTTP status
code to 404. This is useful in telling a caller that
the resources they are trying to get is not found.
*/
func NotFound(writer http.ResponseWriter, response interface{}) {
	WriteJSON(writer, response, 404)
}

/*
SessionExpired is a convenience method for telling clients that a session has
expired. This sends a SessionExpiredResponse structure and
a status of 401 Unauthorized.
*/
func SessionExpired(writer http.ResponseWriter) {
	WriteJSON(writer, NewSessionExpiredResponse(), 401)
}

/*
Success is a convenience method. This will set the HTTP status
code to 200.
*/
func Success(writer http.ResponseWriter, response interface{}) {
	WriteJSON(writer, response, 200)
}
