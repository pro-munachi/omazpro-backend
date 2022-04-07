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

var productCollection *mongo.Collection = database.OpenCollection(database.Client, "product")


func CreateProduct()gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var product models.Product

		if err := c.BindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "hasError": true})
			defer cancel()
			return
		}

		product.ID = primitive.NewObjectID()
		product.Productid = product.ID.Hex()
		product.Createdat, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		product.Updatedat, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		validationErr := validate.Struct(product)
		if validationErr != nil {
			c.JSON(http.StatusOK, gin.H{"message":validationErr.Error(), "hasError": true})
			defer cancel()
			return
		}

		resultInsertionNumber, insertErr := productCollection.InsertOne(ctx, product)
		if insertErr !=nil {
			msg := fmt.Sprintf("product was not created")
			c.JSON(http.StatusOK, gin.H{"message":msg, "hasError": true})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "data":product, "hasError": false, "insertId": resultInsertionNumber})
	}
}

func GetProducts() gin.HandlerFunc{
	return func(c *gin.Context){
		// if err := helper.CheckUserType(c, "ADMIN"); err != nil {
		// 	c.JSON(http.StatusOK, gin.H{"message":err.Error(), "hasError": true})
		// 	return
		// }
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		
		myOptions := options.Find()
		myOptions.SetSort(bson.M{"$natural":-1})
		result,err := productCollection.Find(ctx,  bson.M{}, myOptions)
		
		defer cancel()
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":"error occured while listing products", "hasError": true})
			return
		}
		var allproducts []bson.M
		if err = result.All(ctx, &allproducts); err!=nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "products":allproducts, "hasError": false})}
}


func GetProduct() gin.HandlerFunc{
	return func(c *gin.Context){
		Productid := c.Param("product_id")

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var product models.Product
		err := productCollection.FindOne(ctx, bson.M{"productid":Productid}).Decode(&product)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "product":product, "hasError": false})
	}
}

func DeleteProduct() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("id")

		// primID, _ :=primitive.ObjectIDFromHex(id)

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		res, err := productCollection.DeleteOne(ctx, bson.M{"productid": id})
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "product":res, "hasError": false})
	}
}