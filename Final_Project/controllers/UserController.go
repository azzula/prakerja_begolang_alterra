package controllers

import (
	"net/http"
	"prakerja_begolang_alterra/Final_Project/configs"
	"prakerja_begolang_alterra/Final_Project/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Index(c echo.Context) error {
	var users []models.User

	result := configs.DB.Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Status: false, Message: "Data not found", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "List data user", Data: users,
	})
}

func Store(c echo.Context) error {
	var users models.User
	c.Bind(&users)

	result := configs.DB.First(&models.User{}, "email = ?", users.Email)
	if result.Error != gorm.ErrRecordNotFound {
		return c.JSON(http.StatusConflict, models.BaseResponse{
			Status: false, Message: "Email not be same, please use another email", Data: nil,
		})
	}

	result = configs.DB.Create(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false, Message: "Create data user failed", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "Create data user success", Data: users,
	})
}

func Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var users models.User
	
	result := configs.DB.First(&users, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Status: false, Message: "Data not found", Data: nil,
		})
	}

	c.Bind(&users)

	result = configs.DB.Save(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false, Message: "Update data user failed", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "Update data user success", Data: users,
	})
}

func Destroy(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var users models.User
	
	result := configs.DB.First(&users, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Status: false, Message: "Data not found", Data: nil,
		})
	}

	c.Bind(&users)

	result = configs.DB.Delete(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false, Message: "Delete data user failed", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "Delete data user success", Data: nil,
	})
}

func Search(c echo.Context) error {
	var users []models.User

	result := configs.DB.Find(&users, "name = ?", c.QueryParam("name"))
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Status: false, Message: "Data not found", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "Data user", Data: users,
	})
}

func Show(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var users []models.User

	result := configs.DB.Find(&users, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Status: false, Message: "Data not found", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "Data user", Data: users,
	})
}
