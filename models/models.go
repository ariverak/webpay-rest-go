package sql

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
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
	godotenv.Load()
	config := Config{
		DbName:   os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
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
	// sqlDB, _ := _db.DB()
	Db = _db
	// defer sqlDB.SetMaxIdleConns(1)
}
