package httphandlers

import (
	"fmt"
	"net/http"
)

func pageNotFoundHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "404: Page not found")
}