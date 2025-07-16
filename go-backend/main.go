package main

import (
	"fmt"
	"time"

	"github.com/abdullahedhiii/go-backend/controllers"
	"github.com/abdullahedhiii/go-backend/database"
	"github.com/abdullahedhiii/go-backend/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "traceparent", "tracestate"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	database.ConnectDatabase()

	router.POST("/login", controllers.Login)
	router.POST("/register", controllers.Register)
	router.GET("/books", middleware.AuthMiddleware(), controllers.GetBooks)
	router.GET("/books/:id", middleware.AuthMiddleware(), controllers.GetBook)
	// middleware.AuthMiddleware(),
	fmt.Println("Server is running on port 8080")
	router.Run(":8080")
}
