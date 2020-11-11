package sql

import (
	"time"

	"gorm.io/datatypes"
)

type Order struct {
	ID              int            `gorm:"primaryKey" json:"id"`
	SessionID       string         `gorm:"size:61;index" json:"sessionId"`
	BuyOrder        string         `gorm:"size:26;index" json:"buyOrder"`
	FormAction      string         `gorm:"size:255" json:"formAction"`
	TokenWs         string         `gorm:"size:255" json:"tokenWs"`
	Status          string         `gorm:"size:45" json:"status"`
	Discount        int            `json:"discount"`
	Total           int            `json:"total"`
	WebPayData      datatypes.JSON `json:"webpayData"`
	BillingAddress  string         `gorm:"size:255" json:"billingAddress"`
	BillingAddress2 string         `gorm:"size:255" json:"billingAddress2"`
	BillingAddress3 string         `gorm:"size:255" json:"billingAddress3"`
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
}

func CreateOrder(order Order) Order {
	Db.Create(&order)
	return order
}

func GetOrders() (orders []Order) {
	Db.Find(&orders)
	return
}
