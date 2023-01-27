package router

import (
	"CRUDAPI/database"
	"CRUDAPI/dataentity"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller interface {
	CreateStudent(w http.ResponseWriter, r *http.Request)
	GetStudent(w http.ResponseWriter, r *http.Request)
	GetStudentByID(w http.ResponseWriter, r *http.Request)
	UpdateStudent(w http.ResponseWriter, r *http.Request)
	DeleteStudent(w http.ResponseWriter, r *http.Request)
}
type controller struct {
}

func Getcontroller() Controller {
	return &controller{}
}
func (s *controller) CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var stu dataentity.Studentdata
	json.NewDecoder(r.Body).Decode(&stu)
	database.Database.Create(&stu)
	json.NewEncoder(w).Encode(stu)
}

func (s *controller) GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var students []dataentity.Studentdata
	database.Database.Find(&students)
	json.NewEncoder(w).Encode(students)
}

func (s *controller) GetStudentByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student dataentity.Studentdata
	json.NewDecoder(r.Body).Decode(&student)
	database.Database.First(&student, mux.Vars(r)["eid"])
	json.NewEncoder(w).Encode(student)
}

func (s *controller) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var upstudent dataentity.Studentdata
	database.Database.First(&upstudent, mux.Vars(r)["eid"])
	json.NewDecoder(r.Body).Decode(&upstudent) //decode request body and update
	database.Database.Save(&upstudent)         // updated data save in DB
	json.NewEncoder(w).Encode(upstudent)
}

func (s *controller) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var delstudent dataentity.Studentdata //create variable delstudent
	json.NewDecoder(r.Body).Decode(&delstudent)
	database.Database.Delete(&delstudent, mux.Vars(r)["id"]) // deletes value matching given conditions. If value contains primary key it is included in the conditions.
	json.NewEncoder(w).Encode("employee is deleted ")
}
