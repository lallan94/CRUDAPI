package main

import (
	"CRUDAPI/database"
	"CRUDAPI/router"
)

func main() {
	database.DataMigration()
	router.HandlerRouting()
}
