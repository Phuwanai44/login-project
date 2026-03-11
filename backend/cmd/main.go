package main

import (
	"log"
	"os"

	"login-api/config"
	"login-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// โหลด .env ก่อน
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect database
	config.ConnectDB()

	r := gin.Default()

	// setup routes
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")

	r.Run(":" + port)
}
