# CRUD API(Microservices)
* Basically REST API is a way of creating web services that use the HTTP protocol to transfer data between a client and a server. In this project,I have created a REST API (Microservices) that performs only CRUD (Create, Read, Update, Delete) operations using a MySQL database over HTTP. This REST API would use the net/http package to create an HTTP server and handle all incoming requests. The server can then use the database/sql package to interact with the MySQL database, allowing it to execute SQL queries and perform the necessary CRUD operations.

Before execution this program, first install fallowing driver.
 *    1. Open terminal and go to directory where i have created CRUDAPI folder
 *    2. Installed "go get -u github.com/gorilla/mux" gorilla mux. Gorilla mux is a powerfull router which implements a request router and dispatcher for matching incoming requests to their respective handler.
 *    3. Installed "go get -u gorm.io/gorm" gorm(object relational mapping) driver which help to create table in DB automaticaly.
 *    4. Installed "go get -u gorm.io/driver/mysql" to stablish mysql database connection.
 *    5. installed "go get github.com/joho/godotenv" which loads environment variable from a .env file
* After successfully install all driver we can execute our program.
* Enter "go run main.go" commond and get a popup window(defender firewall) .And allow access.
* Go to Thunderclient or Postman and perform CRUD operation.


* In thunder client enter URL "http://localhost:8080/studentpost" for creation(post) student data.
* We will perform post,get,update and delete all data with POST method. B/c post method is never cacheable and bookmarked over a senstive data.
* For get data enter "/studentget" after port number.
* For getting only one student data enter "studentgetbyid" and send iD number in body in JSON format.
* For update(PUT) data enter "/studentupdate" and send id number(which you want to update) in body.
* For delet data enter "/studentdelete". and send id number in body format.

# Result

*SS of Post data( id no 3).
*![Post data](https://user-images.githubusercontent.com/110240909/216266779-7c5ad13a-d5bd-476f-b8d1-7888b43f638f.png)

* SS of get all data
![Get all data](https://user-images.githubusercontent.com/110240909/216264489-5f4da279-1afa-4802-87e6-0d10aac5f787.png)

* SS of get by ID (id no2).
* ![Get by id](https://user-images.githubusercontent.com/110240909/216264880-a54c5ad8-cd70-4413-9f27-c74cbc0dff0f.png)

* SS of Updated Name, Id and Email ( id no 2).
* ![Update](https://user-images.githubusercontent.com/110240909/216265298-0c41f385-c2ae-4307-8d20-b144bbceab28.png)

*SS of Deleted data (id no 3)
![Delete](https://user-images.githubusercontent.com/110240909/216265425-bfe50b27-3a58-4d86-89c4-2aac935a13ce.png)


