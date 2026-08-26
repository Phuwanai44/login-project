package controllers

import (
	"context"
	"net/http"
	"time"

	"login-api/config"
	"login-api/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CREATE ORDER
func CreateOrder(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var input models.CreateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(input.Items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "items are required"})
		return
	}
	if input.ShippingAddress.FullName == "" || input.ShippingAddress.Address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "shipping address is required"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Stock check: verify each item against product stock field
	for _, item := range input.Items {
		objID, err := primitive.ObjectIDFromHex(item.ProductID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id: " + item.ProductID})
			return
		}

		var product models.Product
		err = config.ProductCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&product)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "product not found: " + item.ProductID})
			return
		}

		if product.Stock < item.Quantity {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "insufficient stock for product: " + product.Name,
			})
			return
		}
	}

	// Deduct stock and calculate total
	total := 0
	for _, item := range input.Items {
		objID, _ := primitive.ObjectIDFromHex(item.ProductID)
		config.ProductCollection.UpdateOne(ctx,
			bson.M{"_id": objID},
			bson.M{"$inc": bson.M{"stock": -item.Quantity}},
		)
		total += item.Price * item.Quantity
	}

	order := models.Order{
		ID:              primitive.NewObjectID(),
		UserID:          userID.(string),
		Items:           input.Items,
		TotalPrice:      total,
		ShippingAddress: input.ShippingAddress,
		Status:          "pending",
		CreatedAt:       time.Now(),
	}

	_, err := config.OrderCollection.InsertOne(ctx, order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot create order"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "order created", "data": order})
}

// GET MY ORDERS
func GetMyOrders(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := config.OrderCollection.Find(ctx, bson.M{"userId": userID.(string)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot get orders"})
		return
	}

	var orders []models.Order
	if err = cursor.All(ctx, &orders); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot decode orders"})
		return
	}

	if orders == nil {
		orders = []models.Order{}
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

// GET ORDER BY ID
func GetOrderById(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var order models.Order
	err = config.OrderCollection.FindOne(ctx, bson.M{
		"_id":    objID,
		"userId": userID.(string),
	}).Decode(&order)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}
