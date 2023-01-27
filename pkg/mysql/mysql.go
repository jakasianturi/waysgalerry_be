package mysql

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {

	// database config from .env
	var DB_HOST = os.Getenv("MYSQLHOST")
	var DB_USER = os.Getenv("MYSQLUSER")
	var DB_PASSWORD = os.Getenv("MYSQLPASSWORD")
	var DB_DATABASENAME = os.Getenv("MYSQLDATABASE")
	var DB_PORT = os.Getenv("MYSQLPORT")

	// connected to database
	var err error
	// dsn := "DB_USER:DB_PASSWORD@tcp(DB_HOST:DB_PORT)/DB_DATABASENAME?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_DATABASENAME)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Connected to Database")
}
