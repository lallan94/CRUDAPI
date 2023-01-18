package database

import (
	"CRUDAPI/dataentity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

var err error

func DataMigration() {
	username := "root"
	password := "Malhotra@123"
	dbname := "sys"
	Database, err = gorm.Open(mysql.Open(username+":"+password+"@/"+dbname), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Database.AutoMigrate(&dataentity.Studentdata{})
}
