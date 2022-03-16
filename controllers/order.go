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

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "product")


func CreateOrder()gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var order models.Order

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "hasError": true})
			defer cancel()
			return
		}

		order.ID = primitive.NewObjectID()
		order.Orderid = order.ID.Hex()
		order.Createdat, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.Updatedat, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		validationErr := validate.Struct(order)
		if validationErr != nil {
			c.JSON(http.StatusOK, gin.H{"message":validationErr.Error(), "hasError": true})
			defer cancel()
			return
		}

		resultInsertionNumber, insertErr := orderCollection.InsertOne(ctx, order)
		if insertErr !=nil {
			msg := fmt.Sprintf("order was not created")
			c.JSON(http.StatusOK, gin.H{"message":msg, "hasError": true})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "data":order, "hasError": false, "insertId": resultInsertionNumber})
	}
}

func GetOrders() gin.HandlerFunc{
	return func(c *gin.Context){
		// if err := helper.CheckUserType(c, "ADMIN"); err != nil {
		// 	c.JSON(http.StatusOK, gin.H{"message":err.Error(), "hasError": true})
		// 	return
		// }
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		
		myOptions := options.Find()
		myOptions.SetSort(bson.M{"$natural":-1})
		result,err := orderCollection.Find(ctx,  bson.M{}, myOptions)
		
		defer cancel()
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":"error occured while listing orders", "hasError": true})
			return
		}
		var allorders []bson.M
		if err = result.All(ctx, &allorders); err!=nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "orders":allorders, "hasError": false})}
}


func GetOrder() gin.HandlerFunc{
	return func(c *gin.Context){
		Orderid := c.Param("order_id")

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var order models.Order
		err := orderCollection.FindOne(ctx, bson.M{"orderid":Orderid}).Decode(&order)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "order":order, "hasError": false})
	}
}