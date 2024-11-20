package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mygo/model"
)

var Db *gorm.DB

func InitDb() {
	dsn := "root:134482@tcp(127.0.0.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = Db.AutoMigrate(&model.User{})
	if err != nil {
		panic("failed to migrate database")
	}
}
