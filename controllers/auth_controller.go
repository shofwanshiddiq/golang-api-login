package controllers

import (
	"api-integration/models"
	"api-integration/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{DB: db}
}

func (ac *AuthController) Register(c *gin.Context) {
	var user models.User

	// Validasi Input Harus JSON
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// Hash Password
	if err := user.HashPassword(user.Password); err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create to Database
	if err := ac.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	// Generate Token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(200, gin.H{"token": token})
}

func (ac *AuthController) Login(c *gin.Context) {
	var loginRequest models.LoginRequest

	// Validasi Input Harus JSON
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// Get User by Email
	var user models.User
	if err := ac.DB.Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	// Check Password
	if err := user.CheckPassword(loginRequest.Password); err != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate Token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(200, gin.H{"token": token})
}
