package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting(c Controller) {
	r := mux.NewRouter()
	r.HandleFunc("/studentpost", c.CreateStudent).Methods("POST") //HandleFunc registers a new route with a matcher for the URL path.
	r.HandleFunc("/studentget", c.GetStudent).Methods("POST")
	r.HandleFunc("/studentgetbyid", c.GetStudentByID).Methods("POST")
	r.HandleFunc("/studentupdate", c.UpdateStudent).Methods("POST")
	r.HandleFunc("/studentdelete", c.DeleteStudent).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
