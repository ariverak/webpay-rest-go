package v1

import (
	"net/http"
	"webpay-rest-go/models"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	var orders models.Order
	result := models.Db.Find(&orders)
	// models.Order.
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": result,
	})
}
