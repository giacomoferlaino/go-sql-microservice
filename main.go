package main

import (
	"fmt"
	"net/http"

	"github.com/giacomoferlaino/go-sql-microservice/sqlmicroservice"
)

func main() {
	app := sqlmicroservice.NewApp()

	app.Router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello world!")
	})

	http.ListenAndServe(":8080", app.Router)
}
