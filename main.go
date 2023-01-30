package main

import (
	"be_test/controller"
	"be_test/database"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"net/http"
)

func main() {
	database.InitDB()
	gorm := database.DB()
	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()
	//db.Create(&model.User{Email: "testing", Password: "testing", Username: "testing"})
	e := echo.New()

	userRoute := e.Group("/users")
	userRoute.POST("/", controller.CreateUser)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"message": "Hello, World!"})
	})
	e.Logger.Fatal(e.Start(":" + viper.GetString("PORT_API")))
}
