package routes

import (
	controller "omazpro/controllers"

	middleware "omazpro/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){

	incomingRoutes.GET("/users/:userid", controller.GetUser())

	incomingRoutes.Use(middleware.Admin())
	incomingRoutes.GET("users", controller.GetUsers())

}