package config

// this contains the database logic
// mysql connections

import (
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// return a varible db
// mysql://myuser:mypassword@localhost:3306/mydatabase
func Connect() *gorm.DB {
	// https://github.com/go-sql-driver/mysql
	dsn := "myuser:mypassword@tcp(localhost:3306)/mydatabase?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error", err)
	}

	return db;
}