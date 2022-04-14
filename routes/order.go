package routes

import (
	controller "omazpro/controllers"
	middleware "omazpro/middleware"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine){


	incomingRoutes.Use(middleware.Authenticate())

	incomingRoutes.POST("order/create", controller.CreateOrder())
	incomingRoutes.GET("/orders/:order_id", controller.GetOrder())
	incomingRoutes.POST("/orders/paid/:id", controller.UpdateOrder())
	incomingRoutes.GET("/orders/user/:id", controller.GetUserOrders())

	incomingRoutes.Use(middleware.Admin())

	incomingRoutes.GET("orders", controller.GetOrders())
	incomingRoutes.POST("/orders/confirm/:id", controller.ConfirmPayment())

}