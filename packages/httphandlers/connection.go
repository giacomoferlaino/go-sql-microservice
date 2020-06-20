package httphandlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/giacomoferlaino/go-sql-microservice/packages/database"
)

// NewConnectionHandler allocates and returns a new ConnectionHandler.
func NewConnectionHandler(appState AppState) *ConnectionHandler {
	return &ConnectionHandler{
		appState: appState,
	}
}

// ConnectionHandler maganges the avaiable database connections.
type ConnectionHandler struct {
	appState AppState
}

// ServeHTTP is the HTTP handler function for this handler.
func (h *ConnectionHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		h.get(res, req)
	case "POST":
		h.post(res, req)
	default:
		pageNotFoundHandler(res, req)
	}
}

func (h *ConnectionHandler) get(res http.ResponseWriter, req *http.Request) {
	bsData, err := json.Marshal(h.appState.Databases())
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}
	fmt.Fprintln(res, string(bsData))
}

func (h *ConnectionHandler) post(res http.ResponseWriter, req *http.Request) {
	reqBody, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(res, err)
		return
	}
	connectionOptions := &database.ConnectionOptions{}
	err = json.Unmarshal(reqBody, connectionOptions)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(res, err)
		return
	}
	dbConnection, err := sql.Open(connectionOptions.Driver, connectionOptions.ConnectionString())
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	h.appState.Databases().Store(connectionOptions.Name, dbConnection)
	fmt.Fprintln(res, "Created connection: ", connectionOptions.Name)
}
