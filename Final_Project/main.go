package main

import (
	"os"
	"final_project/config"
	"final_project/routes"
)

func main() {
	config.Connections()
	e := routes.Route()

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "8080"
	}

	e.Start(":" + port)
	// e.Start(":8080")
}
