package main

import (
	"log"
	"os"

	cors "omazpro/middleware"
	routes "omazpro/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){	

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	if port==""{
		port="8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.CORSMiddleware())

	routes.AuthRoutes(router)
	routes.ProductRoutes(router)
	routes.OrderRoutes(router)
	routes.GalleryRoutes(router)
	routes.UserRoutes(router)

	
	router.GET("/api-1", func(c *gin.Context){
		c.JSON(200, gin.H{"success":"Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context){
		c.JSON(200, gin.H{"success":"Access granted for api-2"})
	})

	router.Run(":" + port)

}	