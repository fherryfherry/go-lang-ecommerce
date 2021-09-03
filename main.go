package main

import (
	"ecommerce/api"
	"ecommerce/database"
	"ecommerce/middelwares"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode("debug") //debug or release

	// Auto migration for the first time
	database.AutoMigrate()

	router := gin.Default()

	// Set max upload limit
	router.MaxMultipartMemory = 8 << 20

	// Auth Routers
	router.POST("/auth/request-token", api.AuthRequestToken)
	router.POST("/auth/refresh-token", api.AuthRefreshToken)

	router.POST("/auth/login", middelwares.ApiGuardMiddleware(), api.AuthLogin)
	router.POST("/auth/logout", middelwares.ApiGuardMiddleware(), api.AuthLogout)

	// Category Routers
	router.GET("/category/list", middelwares.ApiGuardMiddleware(), api.CategoryList)

	// Product Routers
	router.GET("/product/list", middelwares.ApiGuardMiddleware(), api.ProductList)

	// User Routers
	router.POST("/user/update-profile", middelwares.ApiGuardMiddleware(), api.UserUpdateProfile)

	fmt.Println("Application is ready!")

	// Run the server
	router.Run(":8080")
}
