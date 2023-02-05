package main

import (
	"be_test/controller"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"net/http"
)

func main() {

	//db.Create(&model.User{Email: "testing", Password: "testing", Username: "testing"})
	e := echo.New()

	userRoute := e.Group("/users")
	userRoute.POST("/", controller.CreateUser)
	userRoute.GET("/", controller.GetUsers)
	userRoute.GET("/id/:id", controller.GetUserById)
	userRoute.GET("/email/:email", controller.GetUserByEmail)
	userRoute.GET("/username/:username", controller.GetUserByUsername)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"message": "Hello, World!"})
	})
	e.Logger.Fatal(e.Start(":" + viper.GetString("PORT_API")))
}
