package routes

import (
	"final_project/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/users", controllers.Index)
	e.POST("/users", controllers.Store)
	e.PUT("/users/:id", controllers.Update)
	e.DELETE("/users/:id", controllers.Destroy)
	e.GET("/users/q", controllers.Search)
	e.GET("/users/:id", controllers.Show)

	return e
}
