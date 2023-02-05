package database

import (
	"be_test/model"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	dsn := "host=" + viper.GetString("HOST") + " user=" + viper.GetString("USER") + " password=" + viper.GetString("PASSWORD") + " dbname=" + viper.GetString("DBNAME") + " port=" + viper.GetString("PORT") + " sslmode=disable TimeZone=Asia/Shanghai"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	database.AutoMigrate(&model.User{})
	return database
}
