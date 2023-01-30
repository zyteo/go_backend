package database

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func InitDB() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	dsn := "host=" + viper.GetString("HOST") + " user=" + viper.GetString("USER") + " password=" + viper.GetString("PASSWORD") + " dbname=" + viper.GetString("DBNAME") + " port=" + viper.GetString("PORT") + " sslmode=disable TimeZone=Asia/Shanghai"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

}

func DB() *gorm.DB {
	return database
}
