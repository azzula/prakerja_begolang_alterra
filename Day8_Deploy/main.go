package main

import (
	"prakerja_begolang_alterra/Day7_MVC/routes"
	"prakerja_begolang_alterra/Day7_MVC/configs"
)

func main() {
	configs.ConnectDatabase()
	e := routes.InitRoute()
	e.Start(":8000")
}
