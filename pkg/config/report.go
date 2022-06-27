package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Database Connecton Requirements
const (
	dbUserName   = "root"
	dbPassword   = "secret"
	dbHost       = "localhost"
	dbname       = "nmapdojo"
	confCharset  = "utf8mb4"
	confLocation = "Asia%2fTehran"
)

var (
	db *gorm.DB
)

//Connect to database

func Connect() {
	databaseDSN := dbUserName + ":" + dbPassword + "@tcp(" + dbHost + ")/" + dbname + "?charset=" + confCharset + "&parseTime=True&loc=" + confLocation + ""
	connection, err := gorm.Open(mysql.Open(databaseDSN), &gorm.Config{})
	if err != nil {
		panic("Database connection error: " + err.Error())
	}
	db = connection
}

func GetDB() *gorm.DB {
	return db
}
