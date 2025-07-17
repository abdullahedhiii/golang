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
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Request-ID", "Authorization", "traceparent", "tracestate"},
		ExposeHeaders:    []string{"Request-ID"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	database.ConnectDatabase()

	router.POST("/login", middleware.RequestMiddleware(), controllers.Login)
	router.POST("/register", middleware.RequestMiddleware(), controllers.Register)
	router.GET("/books", middleware.RequestMiddleware(), middleware.AuthMiddleware(), controllers.GetBooks)
	router.GET("/books/:id", middleware.RequestMiddleware(), middleware.AuthMiddleware(), controllers.GetBook)
	// middleware.AuthMiddleware(),
	fmt.Println("Server is running on port 8080")
	router.Run(":8080")
}
