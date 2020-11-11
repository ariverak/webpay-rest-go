package sql

import (
	"time"

	"gorm.io/datatypes"
)

type Order struct {
	ID              uint
	SessionID       string `gorm:"size:14"`
	BuyOrder        string `gorm:"size:12"`
	FormAction      string
	TokenWs         string
	Status          string
	Discount        uint
	Total           uint
	WebPayData      datatypes.JSON
	BillingAddress  string
	BillingAddress2 string
	BillingAddress3 string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func CreateOrder(order Order) Order {
	Db.Create(&order)
	return order
}

func GetOrders() (orders []Order) {
	Db.Find(&orders)
	return
}
