package config

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

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, database)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Init database gagal")
	}
	databaseSeeder()
}

func connect() {
	err := godotenv.Load()
	if err != nil {
		panic("Init .env gagal")
	}
}

func databaseSeeder() {
	DB.AutoMigrate(&models.User{})
}
