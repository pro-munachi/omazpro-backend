package routes

import (
	controller "omazpro/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("order/create", controller.CreateOrder())
	incomingRoutes.GET("orders", controller.GetOrders())
	incomingRoutes.GET("/orders/:order_id", controller.GetOrder())

}