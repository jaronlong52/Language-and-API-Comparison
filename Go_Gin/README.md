# Usage Instructions

### Download and Install Go

- https://go.dev/dl/

### Initialize Go Module

`go mod init github.com/<yourusername>/my-gin-api`

### Install Gin and other dependencies

- `go get -u github.com/gin-gonic/gin`
- `go get github.com/joho/godotenv`
- `go get github.com/go-sql-driver/mysql`

### Set up .env file with the following

- DB_HOST
- DB_USER
- DB_PASSWORD
- DB_NAME
- DB_PORT
- PORT (for web access)

### Setup a mySQL database using the script

- SQL_database_script.sql
- Make sure to use `create database <databaseName>` before running the rest of the script.

### Run API

`go run main.go`
