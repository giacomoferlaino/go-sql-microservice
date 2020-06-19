package httphandlers

import (
	"fmt"
	"io"
	"net/http"
)

// NewConnectionHandler allocates and returns a new ConnectionHandler.
func NewConnectionHandler() *ConnectionHandler {
	return &ConnectionHandler{
		write: fmt.Fprint,
	}
}

// ConnectionHandler maganges the avaiable database connections.
type ConnectionHandler struct {
	write func(writer io.Writer, a ...interface{}) (length int, err error)
}

// ServeHTTP is the HTTP handler function for this handler.
func (h *ConnectionHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		h.get(res, req)
	default:
		pageNotFoundHandler(res, req)
	}
}

func (h *ConnectionHandler) get(res http.ResponseWriter, req *http.Request) {
	h.write(res, "GET: Hello World")
}
