package routes

import (
	controller "omazpro/controllers"

	middleware "omazpro/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("products", controller.GetProducts())
	incomingRoutes.GET("/products/:product_id", controller.GetProduct())

	incomingRoutes.Use(middleware.Admin())
	incomingRoutes.POST("products/create", controller.CreateProduct())
	incomingRoutes.DELETE("/products/delete/:id", controller.DeleteProduct())

}