package main

import (
	"be_test/controller"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	e := echo.New()
	userRoute := e.Group("/users")
	userRoute.POST("/", controller.CreateUser)
	userRoute.GET("/", controller.GetUsers)
	userRoute.GET("/id/:id", controller.GetUserById)
	userRoute.GET("/email/:email", controller.GetUserByEmail)
	userRoute.GET("/username/:username", controller.GetUserByUsername)

	e.Logger.Fatal(e.Start(":" + viper.GetString("PORT_API")))
}
