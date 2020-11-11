package v1

import (
	"fmt"
	"net/http"
	"time"
	sql "webpay-rest-go/models"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {

	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	order := sql.Order{
		SessionID: timestamp,
		BuyOrder:  timestamp,
		Status:    "pending",
	}
	sql.CreateOrder(order)
	// models.Order.
	c.JSON(http.StatusOK, gin.H{
		"data": order,
	})
}

func GetOrders(c *gin.Context) {
	orders := sql.GetOrders()
	// models.Order.
	c.JSON(http.StatusOK, gin.H{
		"data": orders,
	})
}
