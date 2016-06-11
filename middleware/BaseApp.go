package middleware

import "github.com/adampresley/webframework/logging"

/*
BaseApp holds context data for the application. This can hold information
such as a database connection, session data, user info, and more. Your middlewares
should attach functions to this structure to pass critical data to request
handlers. Ideally you would create your own application structure and embed this.
*/
type BaseApp struct {
	Log *logging.Logger
}
