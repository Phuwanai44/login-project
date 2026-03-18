package routes

import (
	"login-api/controllers"
	"login-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	//users
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// protected routes
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/profile", controllers.GetProfile)
		auth.PUT("/profile", controllers.UpdateProfile)
	}

	// admin routes
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.AdminOnly())
	{
		admin.GET("/users", controllers.GetUsers)
		admin.DELETE("/users/:id", controllers.DeleteUser) //
	}

	r.GET("/products", controllers.GetProducts)
	r.POST("/products", controllers.CreateProduct)
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)
}
