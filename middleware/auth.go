package middleware

import (
	"fmt"
	"net/http"

	helper "omazpro/helpers"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc{
	return func(c *gin.Context){
		clientToken := c.Request.Header.Get("token")
		if clientToken == ""{
			c.JSON(http.StatusOK, gin.H{"message":fmt.Sprintf("No Authorization header provided"), "hasError": true})
			c.Abort()
			return
		}

		claims, err := helper.ValidateToken(clientToken)
		if err !="" {
			c.JSON(http.StatusOK, gin.H{"message": "token expired", "hasError": true})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("uid",claims.Uid)
		c.Set("user_type", claims.User_type)
		c.Next()
	}
}

func Admin() gin.HandlerFunc{
	return func(c *gin.Context){
		clientToken := c.Request.Header.Get("token")
		if clientToken == ""{
			c.JSON(http.StatusOK, gin.H{"message":fmt.Sprintf("No Authorization header provided"), "hasError": true})
			c.Abort()
			return
		}

		claims, err := helper.ValidateToken(clientToken)
		if err !="" {
			c.JSON(http.StatusOK, gin.H{"message": "token expired", "hasError": true})
			c.Abort()
			return
		}

		if claims.User_type != "ADMIN" {
			c.JSON(http.StatusOK, gin.H{"message": "user is not an admin", "hasError": true})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("uid",claims.Uid)
		c.Set("user_type", claims.User_type)
		c.Next()
	}
}