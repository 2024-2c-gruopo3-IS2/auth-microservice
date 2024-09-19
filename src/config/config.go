// config/config.go
package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

var DB *sqlx.DB
var err error

func InitDB() {

	godotenv.Load()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("SSL_MODE"),
	)
	
	for retries := 5; retries > 0; retries-- {
		fmt.Println("DNS: ", dsn)
		DB, err = sqlx.Connect("postgres", dsn)
        if err == nil {
            if err = DB.Ping(); err == nil {
                log.Println("Successfully connected to the database")
                break
            }
        }

        log.Printf("Error connecting to the database: %v. Retrying in 5 seconds...", err)
        time.Sleep(1 * time.Second)
    }

	createTables()
}

func createTables() {
	userTable := `
	CREATE TABLE IF NOT EXISTS users (
		email VARCHAR(255) PRIMARY KEY,
		password VARCHAR(255) NOT NULL,
		is_blocked BOOLEAN DEFAULT FALSE
	);`

	adminTable := `
	CREATE TABLE IF NOT EXISTS admins (
		email VARCHAR(255) PRIMARY KEY,
		password VARCHAR(255) NOT NULL
	);`

	_, err := DB.Exec(userTable)
	if err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}

	_, err = DB.Exec(adminTable)
	if err != nil {
		log.Fatalf("Failed to create admins table: %v", err)
	}

	fmt.Println("Tables are created or already exist")
}