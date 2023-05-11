package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"prakerja_begolang_alterra/Day8_Deploy/configs"
	"prakerja_begolang_alterra/Day8_Deploy/models"
)

func GetUserController(c echo.Context) error {
	var users []models.User

	result := configs.DB.Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false, Message: "get failed", Data: nil,
		})
	}

	return c.JSON(200, models.BaseResponse{
		Status: true, Message: "success", Data: users,
	})
}

func FindUserController(c echo.Context) error {
	var users []models.User

	result := configs.DB.Find(&users, "name LIKE ?", c.QueryParam("name"))
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false, Message: "get failed", Data: nil,
		})
	}

	return c.JSON(200, models.BaseResponse{
		Status: true, Message: "success", Data: users,
	})
}

func InsertUserController(c echo.Context) error {
	var insertUser models.User
	c.Bind(&insertUser)

	result := configs.DB.First(&models.User{}, "email LIKE ?", insertUser.Email)
	if result.Error != gorm.ErrRecordNotFound {
		return c.JSON(http.StatusConflict, models.BaseResponse{
			Status: false, Message: "email not be same", Data: nil,
		})
	}

	result = configs.DB.Create(&insertUser)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false, Message: "insert failed", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "insert success", Data: insertUser,
	})
}
