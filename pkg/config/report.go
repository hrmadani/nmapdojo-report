package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Database Connecton Requirements
const (
	dbUserName         = "root"
	dbPassword         = "secret"
	dbHost             = "host.docker.internal"
	dbname             = "nmapdojo"
	confCharset        = "utf8mb4"
	confLocation       = "Asia%2fTehran"
	MaxIdleConnections = 20
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
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("error getting database: %v", err)
	}
	sqlDB.SetMaxIdleConns(MaxIdleConnections)
	return db
}
