package main

import (
	"net/http"

	"github.com/DavidQRy/task-manager-gorm/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3020", r)
}
