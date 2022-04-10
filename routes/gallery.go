package routes

import (
	controller "omazpro/controllers"
	middleware "omazpro/middleware"

	"github.com/gin-gonic/gin"
)

func GalleryRoutes(incomingRoutes *gin.Engine){



	incomingRoutes.Use(middleware.Admin())

	incomingRoutes.POST("/gallery/create", controller.CreateGallery())
	incomingRoutes.GET("/gallery/all", controller.GetGallery())
	incomingRoutes.DELETE("/gallery/delete/:id", controller.DeleteGallery())


}