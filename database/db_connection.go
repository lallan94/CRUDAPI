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

type Config struct {
	DB *DBConfig
}
type DBConfig struct {
	Connection string
	Host       string
	Port       string
	Username   string
	Password   string
	Name       string
	Charset    string
}

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GetConfig() *Config {
	loadEnv()
	dbConnection := os.Getenv("DB_CONNECTION")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	return &Config{
		DB: &DBConfig{
			Connection: dbConnection,
			Host:       dbHost,
			Port:       dbPort,
			Username:   dbUsername,
			Password:   dbPassword,
			Name:       dbName,
			Charset:    "utf8",
		},
	}
}

var err error

func DataMigration(config *Config) {
	username := config.DB.Username
	password := config.DB.Password
	dbname := config.DB.Name
	Database, err = gorm.Open(mysql.Open(username+":"+password+"@/"+dbname), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&dataentity.Studentdata{})
}
