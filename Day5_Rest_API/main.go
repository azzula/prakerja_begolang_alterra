package main
import "github.com/labstack/echo/v4"

type User struct {
	ID		int
	Name	string
	Country	string
}

type LoginUser struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func main() {
	e := echo.New()
	e.GET("/hello", HelloController)
	e.GET("/users", UserController)
	e.GET("/users/:id", UserByIDController)
	e.POST("/login", LoginController)

	e.Start(":8000")
}

func HelloController(c echo.Context) error {
	return c.JSON(200, "Hello World")
}

func UserController(c echo.Context) error {
	countryParam := c.QueryParam("country")
	
	user := User{
		ID: 0,
		Name: "Sufyan",
		Country: countryParam,
	}
	return c.JSON(200, user)
}

func UserByIDController(c echo.Context) error {
	id := c.Param("id")

	user := User{
		ID: 0,
		Name: id,
	}
	return c.JSON(200, user)
}

// func LoginController(c echo.Context) error {
// 	email := c.FormValue("email")
// 	password := c.FormValue("password")
// 	user := User{
// 		ID: 0,
// 		Name: email + password,
// 	}
// 	return c.JSON(200, user)
// }

func LoginController(c echo.Context) error {
	loginUser := LoginUser{}
	c.Bind(&loginUser)

	return c.JSON(200, loginUser)
}
