package main

import (
	sql "webpay-rest-go/models"
)

func main() {
	sql.Db.AutoMigrate(&sql.Order{})
	sql.Db.AutoMigrate(&sql.Product{})
	sql.Db.AutoMigrate(&sql.OrderDetail{})
	sql.Db.AutoMigrate(&sql.ProductDetail{})
}
