package router

import (
	"CRUDAPI/database"
	"CRUDAPI/dataentity"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var stu dataentity.Studentdata //create variable"stu" and import Studentdata from dataentity(package) file
	json.NewDecoder(r.Body).Decode(&stu)

	database.Database.Create(&stu) //create a table in databse(package) using .create function in GORM

	json.NewEncoder(w).Encode(stu)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var students []dataentity.Studentdata // create a slice of students

	database.Database.Find(&students) // using .find function in GORM find student data from database
	json.NewEncoder(w).Encode(students)
}

func GetStudentByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var student dataentity.Studentdata // create variable student

	database.Database.First(&student, mux.Vars(r)["eid"]) // First finds the first record ordered by primary key, matching given conditions in DB
	json.NewEncoder(w).Encode(student)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")

	var upstudent dataentity.Studentdata //create variable upstudent

	database.Database.First(&upstudent, mux.Vars(r)["eid"]) //First finds the first record ordered by primary key, matching given conditions in DB

	json.NewDecoder(r.Body).Decode(&upstudent) //decode request body and update

	database.Database.Save(&upstudent) // updated data save in DB
	json.NewEncoder(w).Encode(upstudent)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var delstudent dataentity.Studentdata //create variable delstudent

	database.Database.Delete(&delstudent, mux.Vars(r)["id"]) // deletes value matching given conditions. If value contains primary key it is included in the conditions.
	json.NewEncoder(w).Encode("employee is deleted ")
}
