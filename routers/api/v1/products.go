package v1

import (
	"net/http"
	sql "webpay-rest-go/models"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	products := sql.GetProducts()
	// models.Order.
	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}
