package router

import (
	"CRUDAPI/database"
	"CRUDAPI/dataentity"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Intitalize Interface
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
func (c *controller) CreateStudent(w http.ResponseWriter, r *http.Request) { // initialize methode
	w.Header().Set("Content-Type", "application/json") // By setting the Content-Type header, the server can inform the client that JSON data is being sent. Through the ResponseWriter
	var stu dataentity.Studentdata                     // Create a variable "stu" type of stu is Studentdata which is define in dataentity file.
	json.NewDecoder(r.Body).Decode(&stu)               //The request body is parsed by json NewDecoder and the data is passing the stu struct.
	database.Database.Create(&stu)                     // create a table in database using (gorm.DB).create package.
	json.NewEncoder(w).Encode(stu)
}

func (c *controller) GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var students []dataentity.Studentdata //we have create a slice which will fetch all the data.
	database.Database.Find(&students)     // find all data Using (gorm.DB).Find
	json.NewEncoder(w).Encode(students)   //
}

func (c *controller) GetStudentByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student dataentity.Studentdata
	json.NewDecoder(r.Body).Decode(&student)                        //The request body is parsed by json( like retrive id from request body) an passed to student
	result := database.Database.First(&student, mux.Vars(r)["eid"]) // First finds the first record ordered by primary key, matching given conditions conds.
	if result.Error != nil {
		json.NewEncoder(w).Encode("No data found")
		return
	} else {
		json.NewEncoder(w).Encode(student)
	}

}

func (c *controller) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var upstudent dataentity.Studentdata
	num := mux.Vars(r)["eid"]
	database.Database.First(&upstudent, num)
	json.NewDecoder(r.Body).Decode(&upstudent)
	database.Database.Save(&upstudent)
	json.NewEncoder(w).Encode(&upstudent)
}

func (c *controller) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var delstudent dataentity.Studentdata
	json.NewDecoder(r.Body).Decode(&delstudent)
	result := database.Database.First(&delstudent, mux.Vars(r)["eid"])
	if result.Error != nil {
		json.NewEncoder(w).Encode("This Id data is not available")
		return
	} else {
		database.Database.Delete(&delstudent, mux.Vars(r)["eid"]) // deletes value matching given conditions. If value contains primary key it is included in the conditions.
		json.NewEncoder(w).Encode("employee is deleted ")
	}

}
