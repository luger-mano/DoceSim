package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root_docker"
	password = "steve3571"
	dbname   = "docesim_db"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)
	DB, err = gorm.Open(mysql.Open(mysqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
