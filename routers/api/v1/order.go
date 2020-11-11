package v1

import (
	"fmt"
	"net/http"
	"time"
	sql "webpay-rest-go/models"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {

	order := sql.Order{
		SessionID: fmt.Sprintf("session-%d", time.Now().Unix()),
		// BuyOrder:  timestamp,
		Status: "pending",
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
