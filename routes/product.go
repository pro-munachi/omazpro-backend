package routes

import (
	controller "omazpro/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("products/create", controller.CreateProduct())
	incomingRoutes.GET("products", controller.GetProducts())
	incomingRoutes.GET("/products/:product_id", controller.GetProduct())

}