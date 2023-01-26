package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting(c Controller) {
	r := mux.NewRouter()
	r.HandleFunc("/studentpost", c.CreateStudent).Methods("POST") //HandleFunc registers a new route with a matcher for the URL path.
	r.HandleFunc("/studentget", c.GetStudent).Methods("GET")
	r.HandleFunc("/studentget/{eid}", c.GetStudentByID).Methods("GET")
	r.HandleFunc("/studentupdate/{eid}", c.UpdateStudent).Methods("POST")
	r.HandleFunc("/studentdelete/{id}", c.DeleteStudent).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
