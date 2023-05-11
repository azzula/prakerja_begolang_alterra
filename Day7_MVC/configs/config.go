package configs

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"prakerja_begolang_alterra/Day7_MVC/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/prakerja_begolang_alterra?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Init db failed")
	}
	migrateDatabase()
}

func migrateDatabase() {
	DB.AutoMigrate(&models.User{})
}

// func migrateManualDatabase() {
// }

// func migrateDatabaseV1() {
// 	DB.Migrator().CreateTable(&User{})
// 	DB.Migrator().AddColumn(&User{}, "Name")
// 	DB.Exec("CREATE TABLE ")
// }
