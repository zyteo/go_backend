package controller

import (
	"be_test/database"
	"be_test/logger"
	"be_test/model"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

func CreateUser(c echo.Context) error {
	log := logger.InitLogger()

	u := new(model.User)
	db := database.InitDB()

	if err := c.Bind(u); err != nil {
		data := map[string]interface{}{
			"message": "Failed to create user",
			"error":   err,
		}
		log.Error().Msg("Failed to create user")
		return c.JSON(http.StatusBadRequest, data)
	}

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		data := map[string]interface{}{
			"message": "Failed to hash password",
			"error":   err,
		}
		log.Error().Msg("Failed to hash password")
		return c.JSON(http.StatusBadRequest, data)
	}

	user := &model.User{
		Email:    u.Email,
		Password: string(passwordHashed),
		Username: u.Username,
	}
	if err := db.Create(user).Error; err != nil {
		data := map[string]interface{}{
			"message": "Failed to create user",
			"error":   err,
		}
		log.Error().Msg("Failed to create user")
		return c.JSON(http.StatusBadRequest, data)
	}

	log.Info().Str("email", user.Email).Str("username", user.Username).Msg("User created")
	response := map[string]interface{}{
		"message": "Successfully created user",
		"data":    user,
	}
	return c.JSON(http.StatusCreated, response)
}

func GetUsers(c echo.Context) error {
	log := logger.InitLogger()
	db := database.InitDB()
	var users []model.User
	db.Find(&users)
	log.Info().Msg("Get all users")
	return c.JSON(http.StatusOK, users)
}

func GetUserById(c echo.Context) error {
	log := logger.InitLogger()
	db := database.InitDB()
	var user model.User
	db.Where("id = ?", c.Param("id")).First(&user)
	stringUserID := strconv.Itoa(int(user.ID))
	if stringUserID != c.Param("id") {
		data := map[string]interface{}{
			"error": "User not found",
		}
		log.Error().Msg("User not found")
		return c.JSON(http.StatusBadRequest, data)
	}
	log.Info().Str("email", user.Email).Str("username", user.Username).Msg("Get user by id")
	return c.JSON(http.StatusOK, user)
}

func GetUserByEmail(c echo.Context) error {
	log := logger.InitLogger()
	db := database.InitDB()
	var user model.User
	db.Where("email = ?", c.Param("email")).First(&user)
	if user.Email != c.Param("email") {
		data := map[string]interface{}{
			"error": "User not found",
		}
		log.Error().Msg("User not found")
		return c.JSON(http.StatusBadRequest, data)
	}
	log.Info().Str("email", user.Email).Str("username", user.Username).Msg("Get user by email")
	return c.JSON(http.StatusOK, user)
}

func GetUserByUsername(c echo.Context) error {
	log := logger.InitLogger()
	db := database.InitDB()
	var user model.User
	db.Where("username = ?", c.Param("username")).First(&user)
	if user.Username != c.Param("username") {
		data := map[string]interface{}{
			"error": "User not found",
		}
		log.Error().Msg("User not found")
		return c.JSON(http.StatusBadRequest, data)
	}
	log.Info().Str("email", user.Email).Str("username", user.Username).Msg("Get user by username")
	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	log := logger.InitLogger()
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
		log.Error().Msg("User not found, unable to update")
		return c.JSON(http.StatusBadRequest, data)
	}

	//	get the values of the new email, username and password
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		data := map[string]interface{}{
			"message": "Failed to update user",
			"error":   err,
		}
		log.Error().Msg("Failed to update user")
		return c.JSON(http.StatusBadRequest, data)
	}

	//but if email or username already exists, return error

	if u.Email != user.Email {
		if db.Where("email = ?", u.Email).First(&userEmail).RowsAffected != 0 {
			data := map[string]interface{}{
				"error": "Email already exists",
			}
			log.Error().Msg("Email already exists")
			return c.JSON(http.StatusBadRequest, data)
		}
	}
	if u.Username != user.Username {
		if db.Where("username = ?", u.Username).First(&userUsername).RowsAffected != 0 {
			data := map[string]interface{}{
				"error": "Username already exists",
			}
			log.Error().Msg("Username already exists")
			return c.JSON(http.StatusBadRequest, data)
		}
	}

	//	all ok, update the user
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		data := map[string]interface{}{
			"message": "Failed to hash password",
			"error":   err,
		}
		log.Error().Msg("Failed to hash password")
		return c.JSON(http.StatusBadRequest, data)
	}

	db.Model(&user).Updates(model.User{Email: u.Email, Username: u.Username, Password: string(passwordHashed)})
	response := map[string]interface{}{
		"message": "Successfully updated user",
		"data":    user,
	}
	log.Info().Str("email", user.Email).Str("username", user.Username).Msg("Successfully updated user")
	return c.JSON(http.StatusOK, response)
}

func DeleteUser(c echo.Context) error {
	log := logger.InitLogger()
	db := database.InitDB()

	var user model.User
	db.Where("id = ?", c.Param("id")).First(&user)
	stringUserID := strconv.Itoa(int(user.ID))
	if stringUserID != c.Param("id") {
		data := map[string]interface{}{
			"error": "User not found, unable to delete",
		}
		log.Error().Msg("User not found, unable to delete")
		return c.JSON(http.StatusBadRequest, data)
	}
	db.Delete(&user)
	response := map[string]interface{}{
		"message": "Successfully deleted user",
	}
	log.Info().Str("email", user.Email).Str("username", user.Username).Msg("Successfully deleted user")
	return c.JSON(http.StatusOK, response)

}

func LoginUser(c echo.Context) error {
	log := logger.InitLogger()
	db := database.InitDB()
	var user model.User
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		data := map[string]interface{}{
			"message": "Failed to login user",
			"error":   err,
		}
		log.Error().Msg("Failed to login user")
		return c.JSON(http.StatusBadRequest, data)
	}

	if db.Where("email = ?", u.Email).First(&user).RowsAffected == 0 {
		data := map[string]interface{}{
			"error": "User not found",
		}
		log.Error().Msg("User not found")
		return c.JSON(http.StatusBadRequest, data)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err != nil {
		data := map[string]interface{}{
			"error": "Invalid password",
		}
		log.Error().Msg("Invalid password")
		return c.JSON(http.StatusBadRequest, data)
	}

	response := map[string]interface{}{
		"message": "Successfully logged in",
	}
	log.Info().Str("email", user.Email).Str("username", user.Username).Msg("Successfully logged in")
	return c.JSON(http.StatusOK, response)

}
