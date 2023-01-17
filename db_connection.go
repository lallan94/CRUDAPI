package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var urlDSN = "root:Malhotra@123@tcp(localhost:3306)/sys?parseTime=true"
var err error

func DataMigration() {
	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		panic(err)
		// fmt.Print(err.Error())
		// panic("connection faild :(")
	}
	Database.AutoMigrate(&Studentdata{})
}
