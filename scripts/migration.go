package main

import (
	sql "webpay-rest-go/models"
)

func main() {
	sql.Db.AutoMigrate(&sql.Order{})
}
