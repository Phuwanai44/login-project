package controllers

import (
	"context"
	"login-api/config"
	"login-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

	var users []models.User

	if err = cursor.All(ctx, &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error reading users",
		})
		return
	}

	var result []gin.H

	for _, u := range users {
		result = append(result, gin.H{
			"id":    u.ID.Hex(),
			"name":  u.Name,
			"email": u.Email,
			"role":  u.Role,
		})
	}

	c.JSON(http.StatusOK, result)
}

func GetProfile(c *gin.Context) {

	userID := c.GetString("user_id")

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(400, gin.H{"message": "invalid user id"})
		return
	}

	var user models.User

	err = config.UserCollection.FindOne(
		context.TODO(),
		bson.M{"_id": objID},
	).Decode(&user)

	if err != nil {
		c.JSON(404, gin.H{"message": "user not found"})
		return
	}

	c.JSON(200, gin.H{
		"id":    user.ID.Hex(),
		"name":  user.Name,
		"email": user.Email,
		"role":  user.Role,
	})
}

func UpdateProfile(c *gin.Context) {

	userID := c.GetString("user_id")

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user id"})
		return
	}

	var input struct {
		Name string `json:"name"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "invalid body"})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"name": input.Name,
		},
	}

	result, err := config.UserCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		update,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": "update failed"})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	c.JSON(200, gin.H{
		"message": "profile updated",
	})
}

func DeleteUser(c *gin.Context) {

	id := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	result, err := config.UserCollection.DeleteOne(
		context.TODO(),
		bson.M{"_id": objID},
	)

	if err != nil {
		c.JSON(500, gin.H{"error": "delete failed"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	c.JSON(200, gin.H{
		"message": "user deleted",
	})
}
