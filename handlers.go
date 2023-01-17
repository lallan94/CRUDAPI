package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var stu Studentdata
	json.NewDecoder(r.Body).Decode(&stu)
	Database.Create(&stu)
	json.NewEncoder(w).Encode(stu)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var students []Studentdata
	Database.Find(&students)
	json.NewEncoder(w).Encode(students)
}

func GetStudentByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student Studentdata
	Database.First(&student, mux.Vars(r)["eid"])
	json.NewEncoder(w).Encode(student)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var upstudent Studentdata
	Database.First(&upstudent, mux.Vars(r)["eid"])
	json.NewDecoder(r.Body).Decode(&upstudent)
	Database.Save(&upstudent)
	json.NewEncoder(w).Encode(upstudent)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var delstudent Studentdata
	Database.Delete(&delstudent, mux.Vars(r)["id"])
	json.NewEncoder(w).Encode("employee is deleted ")
}
