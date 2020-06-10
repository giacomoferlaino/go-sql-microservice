package sqlmicroservice

import (
	"fmt"
	"net/http"
)

func main() {
	app := NewApp()

	app.router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello world!")
	})

	http.ListenAndServe(":8080", app.router)
}
