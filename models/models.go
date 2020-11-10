package sql

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Config is the struct for db connection
type Config struct {
	Dialect  string
	Host     string
	Port     string
	User     string
	DbName   string
	Password string
}

//Db is a singleton instance
var Db *gorm.DB

func init() {
	config := Config{
		Dialect:  "mysql",
		DbName:   "ecommerce",
		User:     "root",
		Host:     "127.0.0.1",
		Port:     "3306",
		Password: "a78ma78ma78m",
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DbName,
	)

	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}

	Db = _db
}
