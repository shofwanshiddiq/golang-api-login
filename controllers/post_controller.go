package controllers

import (
	"api-integration/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostController struct {
	DB *gorm.DB
}

func NewPostController(db *gorm.DB) *PostController {
	return &PostController{DB: db}
}

func (pc *PostController) CreateTag(c *gin.Context) {
	var tag models.Tags // Change from 'post' to 'tag' to match the struct

	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	if err := pc.DB.Create(&tag).Error; err != nil { // Use 'tag' instead of undefined 'post'
		c.JSON(500, gin.H{"error": "Failed to create tag"})
		return
	}

	c.JSON(201, gin.H{"data": models.TagResponse{
		ID:   tag.ID, // Use 'tag' instead of undefined 'post'
		Name: tag.Name,
	}})
}

func (pc *PostController) CreatePost(c *gin.Context) {
	var req models.CreatePostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	post := models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userId.(uint),
	}

	tx := pc.DB.Begin()
	if err := tx.Create(&post).Error; err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"error": "Failed to create post"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to commit transaction"})
		return
	}

	// Pengecekan Tag
	if len(req.TagIDs) > 0 {
		var tags []models.Tags

		if err := tx.Find(&tags, req.TagIDs).Error; err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{"error": "Failed to find tags"})
			return
		}

		if len(tags) != len(req.TagIDs) {
			tx.Rollback()
			c.JSON(400, gin.H{"error": "Invalid tag IDs"})
			return
		}

		err := tx.Model(&post).Association("Tags").Append(tags)
		if err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{"error": "Failed to associate tags to post"})
			return
		}

	}

	tx.Commit()

	if err := tx.Commit().Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to commit transaction"})
		return
	}

	c.JSON(201, gin.H{"data": models.PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		Author: models.UserResponse{
			ID:    post.User.ID,
			Name:  post.User.Name,
			Email: post.User.Email,
		},
		Tags: []models.TagResponse{},
	}})
}

func (pc *PostController) GetPosts(c *gin.Context) {
	var posts []models.Post

	// Preload related User and Tags
	if err := pc.DB.Preload("User").Preload("Tags").Find(&posts).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve posts"})
		return
	}

	// Create response slice with the same length as posts
	response := make([]models.PostResponse, len(posts))

	for i, post := range posts {
		response[i] = models.PostResponse{
			ID:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: post.CreatedAt,
			Author: models.UserResponse{
				ID:    post.User.ID,
				Name:  post.User.Name,
				Email: post.User.Email,
			},
			Tags: make([]models.TagResponse, len(post.Tags)), // Correct slice initialization
		}

		// Populate tags inside response
		for j, tag := range post.Tags {
			response[i].Tags[j] = models.TagResponse{
				ID:   tag.ID,
				Name: tag.Name,
			}
		}
	}

	// Return JSON response
	c.JSON(200, gin.H{"data": response})
}

func (pc *PostController) GetPostsByID(c *gin.Context) {
	id := c.Param("id")

	var post models.Post

	if err := pc.DB.Preload("User").Preload("Tags").First(&post, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(200, gin.H{"data": post})
}
