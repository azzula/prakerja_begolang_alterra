package main

import (
	// "os"
	"prakerja_begolang_alterra/Final_Project/config"
	"prakerja_begolang_alterra/Final_Project/routes"
)

func main() {
	config.Connections()
	e := routes.Route()

	// port := os.Getenv("DB_PORT")
	// if port == "" {
	// 	port = "8080"
	// }

	// e.Start(":" + port)
	e.Start(":8080")
}
