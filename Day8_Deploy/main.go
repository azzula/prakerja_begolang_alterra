package main

import (
	"os"
	"prakerja_begolang_alterra/Day8_Deploy/routes"
	"prakerja_begolang_alterra/Day8_Deploy/configs"
)

func main() {
	configs.ConnectDatabase()
	e := routes.InitRoute()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Start(":" + port) // dinamis port
}
