package database

import (
	"github.com/pulsarcoder/learn/reactWithgo/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDatabase() {
	var err error
	Db, err = gorm.Open(sqlite.Open("finData.db"), &gorm.Config{})

	if err != nil {
		panic("faild to connect databse")
	}
	Db.AutoMigrate(&models.User{})

}
