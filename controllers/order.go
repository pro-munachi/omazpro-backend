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

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")


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
		order.User = c.GetString("uid")
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
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "data":order, "id": order.Orderid,  "hasError": false, "insertId": resultInsertionNumber})
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

func UpdateOrder() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("id")
		user := c.GetString("uid")

		var order models.Order
		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
	
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var checkOrder models.Order
		err := orderCollection.FindOne(ctx, bson.M{"orderid":id}).Decode(&checkOrder)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}

		if checkOrder.User != user {
			c.JSON(http.StatusOK, gin.H{"message": "You don't have access to pay for this order", "hasError": true})
			return
		}


		filter := bson.M{"orderid": id}
		set := bson.M{"$set": bson.M{ "IsPaid": order.IsPaid}}
		value, err := orderCollection.UpdateOne(ctx, filter, set)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfullt", "data": value, "order":id, "hasError": false})

	}	
}

func ConfirmPayment() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("id")

		var order models.Order
		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
	
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)


		filter := bson.M{"orderid": id}
		set := bson.M{"$set": bson.M{ "ConfirmPayment": order.ConfirmPayment}}
		value, err := orderCollection.UpdateOne(ctx, filter, set)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfullt", "data": value, "order":id, "hasError": false})

	}	
}


func GetUserOrders() gin.HandlerFunc{
	return func(c *gin.Context){
		user := c.Param("id")

		findOptions := options.Find()


		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var results []models.Order
		curr, err := orderCollection.Find(ctx, bson.M{"user":user}, findOptions)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}

		for curr.Next(context.TODO()) {
			//Create a value into which the single document can be decoded
			var orders models.Order
			err := curr.Decode(&orders)
			if err != nil {
				log.Fatal(err)
			}
	
			results = append(results, orders)
		}

		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "gallery":results, "hasError": false})
	}
}