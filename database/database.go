package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// InitDB initializes the database connection
func InitDB() error {
	var err error

	// Get database URL from environment variable
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		// Default to in-memory SQLite for development
		log.Println("No DATABASE_URL found, using in-memory database")
		return initSQLite()
	}

	// Connect to PostgreSQL
	DB, err = sql.Open("postgres", databaseURL)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// Test the connection
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	log.Println("Successfully connected to PostgreSQL database")

	// Create tables if they don't exist
	if err = createTables(); err != nil {
		return fmt.Errorf("failed to create tables: %v", err)
	}

	return nil
}

// initSQLite initializes an in-memory SQLite database for development
func initSQLite() error {
	var err error
	DB, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		return fmt.Errorf("failed to open SQLite database: %v", err)
	}

	log.Println("Using in-memory SQLite database for development")
	return createTables()
}

// createTables creates all necessary tables
func createTables() error {
	// Users table
	usersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		role VARCHAR(100) NOT NULL,
		avatar TEXT,
		bio TEXT,
		website TEXT,
		location VARCHAR(255),
		skills JSONB,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	// Projects table
	projectsTable := `
	CREATE TABLE IF NOT EXISTS projects (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT NOT NULL,
		tech_stack JSONB,
		status VARCHAR(50) NOT NULL,
		featured BOOLEAN DEFAULT FALSE,
		live_url TEXT,
		github_url TEXT,
		image_url TEXT,
		start_date TIMESTAMP,
		end_date TIMESTAMP,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	// Skills table
	skillsTable := `
	CREATE TABLE IF NOT EXISTS skills (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		category VARCHAR(100) NOT NULL,
		level VARCHAR(50) NOT NULL,
		years_exp INTEGER DEFAULT 0,
		featured BOOLEAN DEFAULT FALSE,
		icon TEXT,
		color VARCHAR(20),
		description TEXT
	);`

	// Contact messages table
	contactTable := `
	CREATE TABLE IF NOT EXISTS contact_messages (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		subject VARCHAR(255) NOT NULL,
		message TEXT NOT NULL,
		status VARCHAR(50) DEFAULT 'unread',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		read_at TIMESTAMP
	);`

	// Visits table for analytics
	visitsTable := `
	CREATE TABLE IF NOT EXISTS visits (
		id SERIAL PRIMARY KEY,
		page VARCHAR(255) NOT NULL,
		user_agent TEXT,
		country VARCHAR(10),
		referrer TEXT,
		ip_address INET,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	tables := []string{usersTable, projectsTable, skillsTable, contactTable, visitsTable}

	for _, table := range tables {
		if _, err := DB.Exec(table); err != nil {
			return fmt.Errorf("failed to create table: %v", err)
		}
	}

	log.Println("Database tables created successfully")
	return nil
}

// Close closes the database connection
func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}