package config

import (
	"../model"
	"github.com/jinzhu/gorm"
)

// DBInit create connection to database
func DBInit() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgres")

	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(model.User{})
	return db
}
