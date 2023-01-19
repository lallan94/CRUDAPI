package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/studentpost", CreateStudent).Methods("POST") //HandleFunc registers a new route with a matcher for the URL path.
	r.HandleFunc("/studentget", GetStudent).Methods("GET")
	r.HandleFunc("/studentget/{eid}", GetStudentByID).Methods("GET")
	r.HandleFunc("/studentput/{eid}", UpdateStudent).Methods("PUT")
	r.HandleFunc("/studentdelete/{id}", DeleteStudent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
