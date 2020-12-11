package bootstrap

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"vulscan/configs"
)

func initDBConnection(conf *configs.Config) *gorm.DB {
	return newConnection("mysql", conf.DBUser, conf.DBPassword, conf.DBPort, conf.DBHost, conf.DBName)
}

func newConnection(dbDriver, dbUser, dbPassword, dbPort, dbHost, dbName string) *gorm.DB {
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser,
		dbPassword, dbHost, dbPort, dbName)

	conn, err := gorm.Open(mysql.Open(DBURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("[Can not connect to database %s]: %s\n", dbDriver, err)
	} else {
		log.Printf("Connected to database: " + dbDriver)
	}
	return conn
}
