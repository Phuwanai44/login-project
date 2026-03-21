package controllers

import (
	"context"
	"math"
	"net/http"
	"time"

	"login-api/config"
	"login-api/models"

	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GET PRODUCTS
func GetProducts(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	search := c.Query("search")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	skip := (page - 1) * limit

	filter := bson.M{}

	if search != "" {
		filter = bson.M{
			"name": bson.M{
				"$regex":   search,
				"$options": "i",
			},
		}
	}

	// ✅ นับจำนวนทั้งหมด
	total, err := config.ProductCollection.CountDocuments(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot count products",
		})
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))

	cursor, err := config.ProductCollection.Find(ctx, filter, findOptions)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot get products",
		})
		return
	}

	var products []models.Product

	if err = cursor.All(ctx, &products); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot decode products",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       products,
		"total":      total,
		"totalPages": totalPages,
		"page":       page,
	})
}

func GetProductById(c *gin.Context) {
	id := c.Param("id")

	// แปลง string → ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var product models.Product

	// หา product ตาม id
	err = config.ProductCollection.FindOne(
		ctx,
		bson.M{"_id": objID},
	).Decode(&product)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "product not found",
		})
		return
	}

	// ส่งข้อมูลกลับ
	c.JSON(http.StatusOK, product)
}

// CREATE PRODUCT
func CreateProduct(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var product models.Product

	// bind json
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// ✅ validation
	if product.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "name is required",
		})
		return
	}

	if product.Price < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "price must be >= 0",
		})
		return
	}

	if product.Stock < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "stock must be >= 0",
		})
		return
	}

	// create id
	product.ID = primitive.NewObjectID()

	// insert
	result, err := config.ProductCollection.InsertOne(ctx, product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot create product",
		})
		return
	}

	// ✅ response กลับไป
	c.JSON(http.StatusCreated, gin.H{
		"message": "product created",
		"id":      result.InsertedID,
		"data":    product,
	})
}

// UPDATE PRODUCT
func UpdateProduct(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})

		return
	}

	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	update := bson.M{
		"$set": bson.M{
			"name":  product.Name,
			"price": product.Price,
			"stock": product.Stock,
		},
	}

	_, err = config.ProductCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot update product",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product updated",
	})
}

// DELETE PRODUCT
func DeleteProduct(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})

		return
	}

	_, err = config.ProductCollection.DeleteOne(ctx, bson.M{"_id": objID})

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot delete product",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product deleted",
	})
}
