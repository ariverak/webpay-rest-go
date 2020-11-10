package v1

import (
	"log"
	"net/http"
	sql "webpay-rest-go/models"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid"
)

func CreateOrder(c *gin.Context) {
	id, err := gonanoid.ID(12)
	if err != nil {
		log.Fatal("Error in generate nanoid")
	}
	order := sql.Order{
		SessionID: "s-" + id,
		BuyOrder:  id,
	}
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
