package database

import (
	"CRUDAPI/dataentity"
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

const projectDirName = "CRUDAPI" // change to relevant project name

type DBConfig struct {
	Connection string
	Host       string
	Port       string
	Username   string
	Password   string
	Name       string
}

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`) // we are loading .env file using loadenv function.

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GetConfig() *DBConfig {
	loadEnv()
	dbConnection := os.Getenv("DB_CONNECTION") // extracting all the data from env file using os.getenv function
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	return &DBConfig{ // it returns  all the configured data in the form of struct
		Connection: dbConnection,
		Host:       dbHost,
		Port:       dbPort,
		Username:   dbUsername,
		Password:   dbPassword,
		Name:       dbName,
	}
}

var err error

func DataMigration(config *DBConfig) {
	username := config.Username
	password := config.Password
	dbname := config.Name
	Database, err = gorm.Open(mysql.Open(username+":"+password+"@/"+dbname), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&dataentity.Studentdata{})
}
