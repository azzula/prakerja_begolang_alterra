package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"prakerja_begolang_alterra/Day8_Deploy/controllers"
)

func InitRoute() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/users", controllers.GetUserController)
	e.GET("/users/q", controllers.FindUserController)
	e.POST("/users", controllers.InsertUserController)
	return e
}
