# CRUDAPI
* Basically REST API is a way of creating web services that use the HTTP protocol to transfer data between a client and a server. In this project,I have created a REST API (Microservices) that performs only CRUD (Create, Read, Update, Delete) operations using a MySQL database over HTTP. This REST API would use the net/http package to create an HTTP server and handle all incoming requests. The server can then use the database/sql package to interact with the MySQL database, allowing it to execute SQL queries and perform the necessary CRUD operations.

Before execution this program, first install fallowing driver.
 *    1. Open terminal and go to directory where i have created CRUDAPI folder
 *    2. Istall "go get -u github.com/gorilla/mux" gorilla mux.
 *    3. Istalled "go get -u gorm.io/gorm" gorm(object relational mapping) driver which help to create table in DB automaticaly.
 *    4. Istalled "go get -u gorm.io/driver/mysql" to stablish mysql database connection.
* After successfully install all driver we can execute our program.
* Enter "go run main.go" commond and get a popup window(defender firewall) .And allow access.
*Go to Thunderclient or Postman and perform CRUD operation.


* In thunder client enter "http://localhost:8080/studentpost" for creation(post) student data.
* For get data enter "/studentget" after port number.
* for update(PUT) data enter "/student/id".
* for delet data enter "/studentdelete/id".
* for getting only one student data enter "studentget/id".

