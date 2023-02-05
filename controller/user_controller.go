package controller

import (
	"be_test/database"
	"be_test/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateUser(c echo.Context) error {
	u := new(model.User)
	db := database.InitDB()

	if err := c.Bind(u); err != nil {
		data := map[string]interface{}{
			"message": "Failed to create user",
			"error":   err,
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	user := &model.User{
		Email:    u.Email,
		Password: u.Password,
		Username: u.Username,
	}
	if err := db.Create(user).Error; err != nil {
		data := map[string]interface{}{
			"message": "Failed to create user",
			"error":   err,
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	response := map[string]interface{}{
		"message": "Successfully created user",
		"data":    user,
	}

	return c.JSON(http.StatusCreated, response)
}

func GetUsers(c echo.Context) error {
	db := database.InitDB()
	var users []model.User
	db.Find(&users)
	return c.JSON(http.StatusOK, users)
}
