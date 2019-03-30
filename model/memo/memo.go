package memo

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Memo struct {
	gorm.Model
	Owner string `json:"owner"`
	Code  string `json:"code"`
	Title string `json:"title"`
	Text  string `json:"text"`
	Tag   string `json:"tag"`
	Time  string `json:"time"`
}

// Connect .
func Connect() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(Memo{})
	return db
}
