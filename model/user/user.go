package user

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	gorm.Model
	Code     string `json:"code"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func Connect() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(User{})
	return db
}
