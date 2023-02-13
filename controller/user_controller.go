package controller

import (
	"be_test/database"
	"be_test/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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

func GetUserById(c echo.Context) error {
	db := database.InitDB()
	var user model.User
	db.Where("id = ?", c.Param("id")).First(&user)
	stringUserID := strconv.Itoa(int(user.ID))
	if stringUserID != c.Param("id") {
		data := map[string]interface{}{
			"error": "User not found",
		}
		return c.JSON(http.StatusBadRequest, data)
	}
	return c.JSON(http.StatusOK, user)
}

func GetUserByEmail(c echo.Context) error {
	db := database.InitDB()
	var user model.User
	db.Where("email = ?", c.Param("email")).First(&user)
	if user.Email != c.Param("email") {
		data := map[string]interface{}{
			"error": "User not found",
		}
		return c.JSON(http.StatusBadRequest, data)
	}
	return c.JSON(http.StatusOK, user)
}

func GetUserByUsername(c echo.Context) error {
	db := database.InitDB()
	var user model.User
	db.Where("username = ?", c.Param("username")).First(&user)
	if user.Username != c.Param("username") {
		data := map[string]interface{}{
			"error": "User not found",
		}
		return c.JSON(http.StatusBadRequest, data)
	}
	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	db := database.InitDB()
	var user model.User
	var userEmail model.User
	var userUsername model.User

	db.Where("id = ?", c.Param("id")).First(&user)

	stringUserID := strconv.Itoa(int(user.ID))
	if stringUserID != c.Param("id") {
		data := map[string]interface{}{
			"error": "User not found, unable to update",
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	//	get the values of the new email, username and password
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		data := map[string]interface{}{
			"message": "Failed to update user",
			"error":   err,
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	//but if email or username already exists, return error

	if u.Email != user.Email {
		if db.Where("email = ?", u.Email).First(&userEmail).RowsAffected != 0 {
			data := map[string]interface{}{
				"error": "Email already exists",
			}
			return c.JSON(http.StatusBadRequest, data)
		}
	}
	if u.Username != user.Username {
		if db.Where("username = ?", u.Username).First(&userUsername).RowsAffected != 0 {
			data := map[string]interface{}{
				"error": "Username already exists",
			}
			return c.JSON(http.StatusBadRequest, data)
		}
	}
	//	all ok, update the user
	db.Model(&user).Updates(model.User{Email: u.Email, Username: u.Username, Password: u.Password})
	response := map[string]interface{}{
		"message": "Successfully updated user",
		"data":    user,
	}
	return c.JSON(http.StatusOK, response)
}

func DeleteUser(c echo.Context) error {
	db := database.InitDB()

	var user model.User
	db.Where("id = ?", c.Param("id")).First(&user)
	stringUserID := strconv.Itoa(int(user.ID))
	if stringUserID != c.Param("id") {
		data := map[string]interface{}{
			"error": "User not found, unable to delete",
		}
		return c.JSON(http.StatusBadRequest, data)
	}
	db.Delete(&user)
	response := map[string]interface{}{
		"message": "Successfully deleted user",
	}
	return c.JSON(http.StatusOK, response)

}
