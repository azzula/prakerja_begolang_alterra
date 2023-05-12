package configs

import (
	"fmt"
	"os"
	"prakerja_begolang_alterra/Final_Project/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connections() {
	connect()

	dbHOST := os.Getenv("DB_HOST")
	dbPORT := os.Getenv("DB_PORT")
	dbUSERNAME := os.Getenv("DB_USERNAME")
	dbPASSWORD := os.Getenv("DB_PASSWORD")
	dbDATABASE := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUSERNAME, dbPASSWORD, dbHOST, dbPORT, dbDATABASE)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Init database failed")
	}
	databaseSeeder()
}

func connect() {
	err := godotenv.Load()
	if err != nil {
		panic("Load .env failed")
	}
}

func databaseSeeder() {
	DB.AutoMigrate(&models.User{})
}
