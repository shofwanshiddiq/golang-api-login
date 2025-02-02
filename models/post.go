package models

import (
	"time"

	"gorm.io/gorm"
)

type Tags struct {
	gorm.Model
	Name  string `json:"name" gorm:"unique"`
	Posts []Post `json:"posts" gorm:"many2many:post_tags"`
}

type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"-"`
	User    User   `json:"author" gorm:"foreignKey:UserID"`
	Tags    []Tags `json:"tags" gorm:"many2many:post_tags"`
}

type PostTag struct {
	PostID uint `gorm:"primaryKey"`
	TagID  uint `gorm:"primaryKey"`
}

// Request Data
type CreatePostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	TagIDs  []uint `json:"tags" binding:"required"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type TagResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type PostResponse struct {
	ID        uint          `json:"id"`
	Title     string        `json:"title"`
	Content   string        `json:"content"`
	CreatedAt time.Time     `json:"created_at"`
	Author    UserResponse  `json:"author"`
	Tags      []TagResponse `json:"tags"`
}
