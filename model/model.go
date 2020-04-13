package model

import (
	"github.com/jinzhu/gorm"
	"github.com/kentio/cmd-health-examination/config"
)

type Hospital struct {
	gorm.Model
	Name    string `gorm: "type:varchar(500)" json:"name"`
	City string		`gorm: "type:varchar(20)" json:"city"`
	Content string `gorm: "type:varchar(1000)" json:"content"`
}


// init db
func InitDB() (*gorm.DB, error)  {
	db, err := gorm.Open("sqlite3", "data.db")
	if err != nil{
		return nil, err
	}
	config.DataDB = db
	//db.AutoMigrate(&City{})
	db.AutoMigrate(&Hospital{})
	db.LogMode(true)
	db.LogMode(true)
	return db, err
}