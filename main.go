package main

import (
	"CRUDAPI/database"
	"CRUDAPI/router"
)

func main() {
	config := database.GetConfig()
	database.DataMigration(config)
	router.HandlerRouting()
}
