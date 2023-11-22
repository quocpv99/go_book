package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	dbUsername := "root"
	dbPassword := "123Aa@"
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "book"
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)
	d, err := gorm.Open("mysql", dbURL)
	if err != nil {
		panic("Failed to connect to the database")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
