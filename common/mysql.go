package common

import (
	"fmt"
	"gin-bubble/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
)

var (
	DB *gorm.DB
)

func InitMySQL() *gorm.DB {
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))

	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Todo{})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
