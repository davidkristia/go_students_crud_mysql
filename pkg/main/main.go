package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/davidkristia/go_students_crud_mysql/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterStudentsRoutes(r)
	http.Handle("/", r)
	fmt.Println("Starting Server localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
