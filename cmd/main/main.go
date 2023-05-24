package main

import (
	"log"
	"net/http"

	"github.com/NileshMutyam/UserManagementSystem/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterUserRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}
