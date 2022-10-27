package databases

import (
	"MyGram/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// config db
const (
	username = "root"
	password = ""
	hostname = "127.0.0.1:3306"
	dbname   = "db-mygram"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {

	// conn db
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local", username, password, hostname, dbname))
	if err != nil {
		panic(err)
	}

	db = db.Debug()
	autoMigrate()
}

func GetDb() *gorm.DB {
	return db
}

func autoMigrate() {
	db.AutoMigrate(&models.User{}, &models.Photo{}, &models.Comments{}, &models.SocialMedia{})
}
