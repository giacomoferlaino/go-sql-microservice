package main

import (
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/giacomoferlaino/go-sql-microservice/packages/sqlmicroservice"
)

func main() {
	app := sqlmicroservice.NewApp()

	app.DefaultHandlers()

	http.ListenAndServe(":8080", app.Router())
}
