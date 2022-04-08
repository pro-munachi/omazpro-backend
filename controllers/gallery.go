package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"omazpro/database"
	"omazpro/models"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var galleryCollection *mongo.Collection = database.OpenCollection(database.Client, "gallery")


func CreateGallery()gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var gallery models.Gallery

		if err := c.BindJSON(&gallery); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "hasError": true})
			defer cancel()
			return
		}

		gallery.ID = primitive.NewObjectID()
		gallery.Picid = gallery.ID.Hex()
		gallery.Createdat, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		gallery.Updatedat, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		validationErr := validate.Struct(gallery)
		if validationErr != nil {
			c.JSON(http.StatusOK, gin.H{"message":validationErr.Error(), "hasError": true})
			defer cancel()
			return
		}

		resultInsertionNumber, insertErr := galleryCollection.InsertOne(ctx, gallery)
		if insertErr !=nil {
			msg := fmt.Sprintf("pic was not created")
			c.JSON(http.StatusOK, gin.H{"message":msg, "hasError": true})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, gin.H{"message": "image uploaded successfully", "data":gallery, "hasError": false, "insertId": resultInsertionNumber})
	}
}

func GetCategory() gin.HandlerFunc{
	return func(c *gin.Context){
		Category := c.Param("category")

		findOptions := options.Find()


		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var results []models.Gallery
		curr, err := galleryCollection.Find(ctx, bson.M{"category":Category}, findOptions)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}

		for curr.Next(context.TODO()) {
			//Create a value into which the single document can be decoded
			var elem models.Gallery
			err := curr.Decode(&elem)
			if err != nil {
				log.Fatal(err)
			}
	
			results = append(results, elem)
		}

		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "gallery":results, "hasError": false})
	}
}


func GetGallery() gin.HandlerFunc{
	return func(c *gin.Context){
		// if err := helper.CheckUserType(c, "ADMIN"); err != nil {
		// 	c.JSON(http.StatusOK, gin.H{"message":err.Error(), "hasError": true})
		// 	return
		// }
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		
		myOptions := options.Find()
		myOptions.SetSort(bson.M{"$natural":-1})
		result,err := galleryCollection.Find(ctx,  bson.M{}, myOptions)
		
		defer cancel()
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":"error occured while listing orders", "hasError": true})
			return
		}
		var allgallery []bson.M
		if err = result.All(ctx, &allgallery); err!=nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "orders":allgallery, "hasError": false})}
}