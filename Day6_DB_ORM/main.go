package main
import "github.com/labstack/echo/v4"

type User struct {
	Id		int			`json:"id"`
	Name	string		`json:"name"`
	Address	string		`json:"address"`
	Email	string		`json:"email"`
}

type BaseResponse struct {
	Status	bool		`json:"status"`
	Message string		`json:"message"`
	Data	interface{}	`json:"data"`
}

func main() {
	e := echo.New()
	e.GET("/users", GetUserController)
	e.Start(":8000")
}

func GetUserController(c echo.Context) error {
	var users []User
	// users = append(users, User{1, "a", "a", "a"})
	// users = append(users, User{2, "b", "b", "b"})

	var insertUser User
	if insertUser.Email != "" {
		
	}
	// insertUser.insert()

	return c.JSON(200, BaseResponse{
		true, "success", "Hello",
	})
}