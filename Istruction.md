# CRUDAPI
//before execution this program, fist install fallowing driver.
    // 1. open terminal and go to directory where i have created CRUDAPI folder
    //2. istall "go get -u github.com/gorilla/mux" gorilla mux.
    //3.istalled "go get -u gorm.io/gorm" gorm(object relational mapping) driver which help to create table in DB automaticaly.
    //4.istalled "go get -u gorm.io/driver/mysql" to stablish mysql database connection.
//after successfully install all driver we can execute our program.
//enter "go run main.go" commond and get a popup window(defender firewall) .And allow access.
//Go to Thunderclient or Postman and perform CRUD operation.


In thunder client enter "http://localhost:8080/studentpost" for creation(post) student data.
For get data enter "/studentget" after port number.
for update(PUT) data enter "/student/id".
for delet data enter "/studentdelete/id".
for getting only one student data enter "studentget/id".
