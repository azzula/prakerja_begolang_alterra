package configs

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"prakerja_begolang_alterra/Day8_Deploy/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	loadENV()

	dbHOST := os.Getenv("DB_HOST")
	dbPORT := os.Getenv("DB_PORT")
	dbUSER := os.Getenv("DB_USER")
	dbPASSWORD := os.Getenv("DB_PASSWORD")
	dbDATABASE := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUSER, dbPASSWORD, dbHOST, dbPORT, dbDATABASE)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Init db failed")
	}
	migrateDatabase()
}

func loadENV() {
	err := godotenv.Load()
	if err != nil {
		panic("Load .env failed")
	}
}

func migrateDatabase() {
	DB.AutoMigrate(&models.User{})
}
