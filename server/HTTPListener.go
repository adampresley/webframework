package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

/*
HTTPListener is a structure which provides an HTTP listener to service
requests. This structure offers methods to add routes and middlewares. Typical
usage would first call NewHTTPListener(), add routes, then call
StartHTTPListener.
*/
type HTTPListener struct {
	Address string
	Port    int

	Router                 *mux.Router
	BaseMiddlewareHandlers alice.Chain
}

/*
NewHTTPListener creates a new instance of the HTTPListener
*/
func NewHTTPListener(address string, port int) *HTTPListener {
	return &HTTPListener{
		Address: address,
		Port:    port,
		Router:  mux.NewRouter(),
	}
}

/*
AddMiddleware adds a new middleware handler to the request chain.
*/
func (listener *HTTPListener) AddMiddleware(middlewareHandler alice.Constructor) *HTTPListener {
	listener.BaseMiddlewareHandlers = listener.BaseMiddlewareHandlers.Append(middlewareHandler)
	return listener
}

/*
AddRoute adds a HTTP handler route to the HTTP listener.
*/
func (listener *HTTPListener) AddRoute(path string, handlerFunc http.HandlerFunc, methods ...string) *HTTPListener {
	listener.Router.Handle(path, listener.BaseMiddlewareHandlers.ThenFunc(handlerFunc)).Methods(methods...)
	return listener
}

/*
AddRouteWithMiddleware adds a HTTP handler route that goes through an additional
middleware handler, to the HTTP listener. This is particularly useful
for setting up routes that require authentication.
*/
func (listener *HTTPListener) AddRouteWithMiddleware(path string, handlerFunc http.HandlerFunc, middlewareHandler alice.Constructor, methods ...string) *HTTPListener {
	listener.Router.Handle(
		path,
		listener.BaseMiddlewareHandlers.Append(middlewareHandler).ThenFunc(handlerFunc),
	).Methods(methods...)

	return listener
}

/*
AddStaticRoute adds a HTTP handler route for static assets. Note that this does not work
with tools that embed assets in your executable. To do that you would need to setup
your own static file handler manually.
*/
func (listener *HTTPListener) AddStaticRoute(pathPrefix string, directory string) *HTTPListener {
	fileServer := http.FileServer(http.Dir(directory))
	listener.Router.PathPrefix(pathPrefix).Handler(http.StripPrefix(pathPrefix, fileServer))

	return listener
}

/*
Start starts the HTTP listener and servicing requests.
*/
func (listener *HTTPListener) Start() error {
	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", listener.Address, listener.Port),
		Handler: alice.New().Then(listener.Router),
	}

	return httpServer.ListenAndServe()
}

/*
StartTLS starts the HTTP listener with SSL and services requests
*/
func (listener *HTTPListener) StartTLS(certificateFile, keyFile string) error {
	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", listener.Address, listener.Port),
		Handler: alice.New().Then(listener.Router),
	}

	return httpServer.ListenAndServeTLS(certificateFile, keyFile)
}
