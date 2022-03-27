package routes

import (
	controller "omazpro/controllers"
	middleware "omazpro/middleware"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine){

	incomingRoutes.Use(middleware.Authenticate())

	incomingRoutes.POST("order/create", controller.CreateOrder())


	incomingRoutes.Use(middleware.Admin())

	incomingRoutes.GET("orders", controller.GetOrders())
	incomingRoutes.GET("/orders/:order_id", controller.GetOrder())

}