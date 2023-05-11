package main
import (
	"net/http"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	Id			int			`json:"id"`
	Name		string		`json:"name"`
	Address		string		`json:"address"`
	Email		string		`json:"email"`
	Password	string		`json:"password"`
}

type BaseResponse struct {
	Status	bool		`json:"status"`
	Message string		`json:"message"`
	Data	interface{}	`json:"data"`
}

func main() {
	connectDatabase()
	e := echo.New()
	e.GET("/users", GetUserController)
	e.GET("/users/q", FindUserController)
	e.POST("/users", InsertUserController)
	e.Start(":8000")
}

func connectDatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/prakerja_begolang_alterra?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
  	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Init db failed")
	}
	migrateDatabase()
}

func migrateDatabase() {
	DB.AutoMigrate(&User{})
}

func GetUserController(c echo.Context) error {
	var users []User
	// users = append(users, User{1, "a", "a", "a"})
	// users = append(users, User{2, "b", "b", "b"})

	result := DB.Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse{
			false, "get failed", nil,
		})
	}

	return c.JSON(200, BaseResponse{
		true, "success", users,
	})
}

func FindUserController(c echo.Context) error {
	var users []User

	result := DB.Find(&users, "name LIKE ?", c.QueryParam("name"))
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse{
			false, "get failed", nil,
		})
	}

	return c.JSON(200, BaseResponse{
		true, "success", users,
	})
}

func InsertUserController(c echo.Context) error {
	var insertUser User
	c.Bind(&insertUser)

	// if insertUser.Email != "" {
	// 	return c.JSON(http.StatusBadRequest, BaseResponse)
	// }
	// insertUser.insert()

	result := DB.Create(&insertUser)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse{
			false, "insert failed", nil,
		})
	}
	return c.JSON(http.StatusOK, BaseResponse{
		true, "insert success", insertUser,
	})
}