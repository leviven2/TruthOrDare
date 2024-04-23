package database

import (
	"TruthOrDare/models"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var database *gorm.DB

func connect() {
	var err error

	if os.Getenv("GO_APP_MODE") == "production" {
		database, err = gorm.Open(mysql.Open(os.Getenv("DB_CONNECT_STRING")), &gorm.Config{})
	} else {
		database, err = gorm.Open(sqlite.Open("database.sqlite"))
	}
	if err != nil {
		panic(err)
	}

	err = database.AutoMigrate(&models.User{}, &models.Category{}, &models.Item{})

	if err != nil {
		panic(err)
	}
}

func DB() *gorm.DB {
	if database == nil {
		connect()
	}
	return database
}
