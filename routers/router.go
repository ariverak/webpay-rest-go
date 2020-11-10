package routers

import (
	"webpay-rest-go/middlewares/cors"
	v1 "webpay-rest-go/routers/api/v1"

	"github.com/gin-gonic/gin"
)

//InitRouter is the function
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(cors.CORS())

	apiv1 := r.Group("/api/v1")
	apiv1.Use()
	{
		//get all orders
		apiv1.GET("/orders", v1.GetOrders)
		apiv1.POST("/orders", v1.CreateOrder)
	}

	return r
}
