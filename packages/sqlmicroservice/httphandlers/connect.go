package httphandlers

import (
	"fmt"
	"net/http"
)

// ConnectHandler implements the handling function for the "/connect" API route
func ConnectHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		getHandler(res, req)
	case "POST":
		postHandler(res, req)
	default:
		pageNotFoundHandler(res, req)
	}
}

func getHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "GET: Hello world!")
}

func postHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "POST: Hello world!")
}
