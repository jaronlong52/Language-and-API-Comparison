package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Build the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Open DB connection
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// Verify DB connection
	if err := db.Ping(); err != nil {
		log.Fatal("Database unreachable:", err)
	}

	fmt.Println("✅ Connected to MySQL")
}

func main() {
	router := gin.Default()

	router.GET("/api", helloHandler)
	router.GET("/api/users", getUsersHandler)
	router.POST("/api/addUser", addUserHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	router.Run("127.0.0.1:" + port)
}

// GET /api
func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}

// GET /api/users
func getUsersHandler(c *gin.Context) {
	rows, err := db.Query("SELECT id, username, name FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var users []map[string]interface{}
	for rows.Next() {
		var id int
		var username, name string
		err := rows.Scan(&id, &username, &name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, gin.H{"id": id, "username": username, "name": name})
	}

	c.JSON(http.StatusOK, users)
}

// POST /api/addUser
func addUserHandler(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Name     string `json:"name"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if input.Username == "" || input.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing username or name"})
		return
	}

	query := "INSERT INTO users (username, name) VALUES (?, ?)"
	_, err := db.Exec(query, input.Username, input.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "User added successfully",
		"username": input.Username,
	})
}
