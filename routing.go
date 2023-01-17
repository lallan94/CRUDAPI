package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/student", CreateStudent).Methods("POST")
	r.HandleFunc("/students", GetStudent).Methods("GET")
	r.HandleFunc("/student/{eid}", GetStudentByID).Methods("GET")
	r.HandleFunc("/student/{eid}", UpdateStudent).Methods("PUT")
	r.HandleFunc("/student/{id}", DeleteStudent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
