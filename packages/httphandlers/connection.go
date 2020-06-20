package httphandlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// NewConnectionHandler allocates and returns a new ConnectionHandler.
func NewConnectionHandler(appState AppState) *ConnectionHandler {
	return &ConnectionHandler{
		appState: appState,
		write:    fmt.Fprint,
	}
}

// ConnectionHandler maganges the avaiable database connections.
type ConnectionHandler struct {
	appState AppState
	write    func(writer io.Writer, a ...interface{}) (length int, err error)
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
	bsData, err := json.Marshal(h.appState.Databases())
	if err != nil {
		h.write(res, "error")
		return
	}
	h.write(res, string(bsData))
}
