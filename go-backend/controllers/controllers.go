package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/abdullahedhiii/go-backend/database"
	"github.com/abdullahedhiii/go-backend/models"
	caching "github.com/abdullahedhiii/go-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var SecretKey = "edhi"

func Login(c *gin.Context) {
	var reqBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	fmt.Println("Attempting to login with :", reqBody.Email, reqBody.Password)

	result := database.DB.Where("email = ?", reqBody.Email).First(&user)
	if result.Error != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}
	fmt.Println("Fetched user from DB:", user.Email, user.Password)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqBody.Password))

	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	token, _ := claims.SignedString([]byte(SecretKey))

	user.Password = ""
	c.JSON(200, gin.H{"message": "Login successful", "token": token, "user": user})
}

func Register(c *gin.Context) {
	var reqBody struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{
		Name:     reqBody.Name,
		Email:    reqBody.Email,
		Password: string(hashedPassword),
	}

	fmt.Println("Registering user:", user.Email)
	fmt.Println("Password hash:", user.Password)

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(201, gin.H{"message": "User registered", "user": user})
}

func GetBooks(c *gin.Context) {
	var books []models.Book

	result := database.DB.Find(&books)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, books)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	res, err := caching.GET("book", id)
	if err == nil && len(res) > 0 {
		fmt.Println("Cache hit for book ID:", id)
		book.ID = id
		book.Title = res["Title"]
		book.Author = res["Author"]
		c.JSON(200, book)
		return
	}

	result := database.DB.First(&book, "id = ?", id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}

	fmt.Println("Cache miss for book ID:", id)
	err = caching.PUT("book", id, book, 10*time.Minute)
	if err != nil {
		fmt.Println("Error caching book:", err)
		c.JSON(404, gin.H{"error": "Error caching book"})

	} else {
		fmt.Println("Book cached successfully for ID:", id)
		c.JSON(200, book)

	}
}
