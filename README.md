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

## Running the API
- Start the API using:
  ```sh
  go run main.go
