package controllers

import (
	"net/http"
	"prakerja_begolang_alterra/Final_Project/config"
	"prakerja_begolang_alterra/Final_Project/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Index(c echo.Context) error {
	var users []models.User

	result := config.DB.Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Status: false, Message: "Data tidak ada", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "List data user", Data: users,
	})
}

func Store(c echo.Context) error {
	var users models.User
	c.Bind(&users)

	result := config.DB.First(&models.User{}, "email = ?", users.Email)
	if result.Error != gorm.ErrRecordNotFound {
		return c.JSON(http.StatusConflict, models.BaseResponse{
			Status: false, Message: "Email telah ditemukan, silakan gunakan email lain", Data: nil,
		})
	}

	result = config.DB.Create(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false, Message: "Gagal tambah data", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "Berhasil tambah data", Data: users,
	})
}

func Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var users models.User
	
	result := config.DB.First(&users, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Status: false, Message: "Data tidak ada", Data: nil,
		})
	}

	c.Bind(&users)

	result = config.DB.Save(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false, Message: "Gagal perbarui data", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "Berhasil perbarui data", Data: users,
	})
}

func Destroy(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var users models.User
	
	result := config.DB.First(&users, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Status: false, Message: "Data tidak ada", Data: nil,
		})
	}

	c.Bind(&users)

	result = config.DB.Delete(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false, Message: "Gagal hapus data", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "Berhasil hapus data", Data: nil,
	})
}

func Search(c echo.Context) error {
	var users []models.User

	result := config.DB.Find(&users, "name = ?", c.QueryParam("name"))
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Status: false, Message: "Data tidak ada", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "Data user", Data: users,
	})
}

func Show(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var users models.User
	
	result := config.DB.First(&users, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Status: false, Message: "Data tidak ada", Data: nil,
		})
	}

	c.Bind(&users)

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "Data user", Data: users,
	})
}
