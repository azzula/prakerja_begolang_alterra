package main

import (
	"os"
	"prakerja_begolang_alterra/Final_Project/configs"
	"prakerja_begolang_alterra/Final_Project/routes"
)

func main() {
	configs.Connections()
	e := routes.Route()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Start(":" + port)
}
