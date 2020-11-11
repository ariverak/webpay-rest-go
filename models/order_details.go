package sql

import (
	"time"
)

type OrderDetail struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Discount  int       `json:"discount"`
	Quantity  int       `json:"quantity"`
	SubTotal  int       `json:"subtotal"`
	OrderID   int       `json:"orderId"`
	Order     Order     `gorm:"foreignKey:OrderID" json:"order"`
	ProductID int       `json:"productId"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func CreateOrderDetail(orderDetail OrderDetail) OrderDetail {
	Db.Create(&orderDetail)
	return orderDetail
}
