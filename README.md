# Golang API User Login Authorization

This is a RESTful API built using Golang with the Gin framework, GORM as the ORM for MySQL, JWT for authentication, and middleware for request handling. The API supports user authentication, tag creation, and post creation



### Features
* User registration with SHA-256 password hashing
* User login with JWT-based authentication
* Middleware for protected routes
* CRUD operations for users, posts, and tags
* Many-to-Many relationship between posts and tags
* GORM-based relational database modeling

# Technologies
![Golang](https://img.shields.io/badge/golang-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)  ![REST API](https://img.shields.io/badge/restapi-%23000000.svg?style=for-the-badge&logo=swagger&logoColor=white)   ![MySQL](https://img.shields.io/badge/mysql-%234479A1.svg?style=for-the-badge&logo=mysql&logoColor=white)  

Uses golang as main frameworks for build an API, with RESTful API for communication with mySQL database

* Golang (Gin-Gonic framework) 
* MySQL (Database)
* GORM (ORM for database operations)
* JWT (JSON Web Token for authentication)
* SHA-256 (Password hashing)

# API Endpoints Documentation

This document provides an overview of the API endpoints, their methods, and functionality.

## Endpoints Table

| Method     | API Endpoint               | Description                                      | Pages             |
|------------|---------------------------|--------------------------------------------------|-------------------|
| **POST**   | `/api/auth/register`       | Register a new user                             | Authentication    |
| **POST**   | `/api/auth/login`          | Authenticate user and return JWT token         | Authentication    |
| **GET**    | `/api/users`               | Fetch list of users                            | Users Page        |
| **POST**   | `/api/users`               | Create a new user                              | Users Page        |
| **POST**   | `/api/users_without_db`    | Create a user without storing in database      | Users Page        |
| **GET**    | `/api/users_without_db`    | Fetch users created without database storage   | Users Page        |
| **POST**   | `/api/tags`                | Create a new tag                               | Posts Page        |
| **POST**   | `/api/posts`               | Create a new post                              | Posts Page        |
| **GET**    | `/api/posts`               | Fetch list of posts                            | Posts Page        |
| **GET**    | `/api/posts/{id}`          | Fetch a specific post by ID                    | Posts Page        |

### Authentication
- The authentication routes (`/api/auth/register` and `/api/auth/login`) handle user registration and authentication.
- The login endpoint returns a JWT token that is required for accessing protected routes.

### Users
- The `/api/users` endpoints allow fetching and creating users.
- The `/api/users_without_db` endpoints are for managing temporary users that are not stored in the database.

### Posts & Tags
- The `/api/posts` endpoints allow for creating and fetching posts.
- The `/api/tags` endpoint is used to create new tags.

### Security & Middleware
- Protected routes require a valid JWT token for access.
- Middleware ensures that unauthorized requests are blocked.

# Database Structure
## User Model
```golang
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"Password" binding:"required"`
}
```

## Post Model
```golang
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
```

# Golang API Starter Guide

This guide will help you set up and run the Golang API using Gin-Gonic, GORM, and MySQL.

## Prerequisites

Ensure you have the following installed:

- [Golang](https://go.dev/dl/) (latest version)
- [MySQL](https://dev.mysql.com/downloads/)
- [Git](https://git-scm.com/)

## Initialization

Follow these steps to set up the project:

### 1. Initialize the Go Module
Run the following command in the project directory:

```sh
go mod init api-integration
```

### 2. Install Dependencies
Install the required packages:

```sh
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u github.com/joho/godotenv
```

### 3. Configure Database
Create a .env file in the root directory and add your database credentials:

```env
DB_USER=root
DB_PASSWORD=password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=dbname

JWT_SECRET_KEY=super_secret_key
JWT_EXPIRATION_IN=24h
```
### 4. Run API
```sh
go run main.go
```
