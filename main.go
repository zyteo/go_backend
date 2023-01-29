package main

import (
	"be_test/model"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	// connect to postgres db with gorm and insert user
	dsn := "host=" + viper.GetString("HOST") + " user=" + viper.GetString("USER") + " password=" + viper.GetString("PASSWORD") + " dbname=" + viper.GetString("DBNAME") + " port=" + viper.GetString("PORT") + " sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&model.User{})

	db.Create(&model.User{Email: "testing", Password: "testing", Username: "testing"})
}
