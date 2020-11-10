package main

import (
	"webpay-rest-go/models"
)

func main() {
	models.Db.AutoMigrate(&models.Order{})
}
