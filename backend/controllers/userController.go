package controllers

import (
	"context"
	"login-api/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Profile(c *gin.Context) {

	email, _ := c.Get("email")
	role, _ := c.Get("role")

	c.JSON(200, gin.H{
		"email": email,
		"role":  role,
	})
}

func GetUsers(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := config.UserCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot get users",
		})
		return
	}

	var users []interface{}

	if err = cursor.All(ctx, &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error reading users",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}
