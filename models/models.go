package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" //Needed for GORM
)

var db gorm.DB

//Setup the models
func Setup() error {
	db, err := gorm.Open("postgres", "user=gorm dbname=gorm")

	db.AutoMigrate(&Node{}, &Reading{})

	db.DB().Ping()
	return err
}
