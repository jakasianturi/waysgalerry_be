package postgresql

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {

	// database config from .env
	var DB_HOST = os.Getenv("DB_HOST")
	var DB_USER = os.Getenv("DB_USER")
	var DB_PASSWORD = os.Getenv("DB_PASSWORD")
	var DB_DATABASENAME = os.Getenv("DB_DATABASENAME")
	var DB_PORT = os.Getenv("DB_PORT")

	// connected to database
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", DB_HOST, DB_USER, DB_PASSWORD, DB_DATABASENAME, DB_PORT)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connected to Database")
}
