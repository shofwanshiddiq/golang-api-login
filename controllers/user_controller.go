package controllers

import (
	"api-integration/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Bikin Array Buat Nampung data user

var usersInMemory = []models.User{}

type UserController struct {
	DB *gorm.DB
}

func GetUserWithoutDB(c *gin.Context) {
	c.JSON(200, gin.H{"data": usersInMemory})
}

func CreateUserWithoutDB(c *gin.Context) {
	var users models.User
	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	users.ID = uint(len(usersInMemory) + 1)
	usersInMemory = append(usersInMemory, users)
	c.JSON(201, gin.H{"data": users})
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

func (uc *UserController) GetUsers(c *gin.Context) {
	var users []models.User
	uc.DB.Find(&users)
	c.JSON(200, gin.H{"data": users})
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var users models.User

	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	result := uc.DB.Create(&users)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(201, gin.H{"data": users})
}
