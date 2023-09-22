package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

var (
	db     *sql.DB
	dbOnce sync.Once
)

// GetDB returns a singleton database connection.
func GetDB() *sql.DB {
	dbOnce.Do(func() {
		// Initialize the database connection
		var err error
		db, err = sql.Open("postgres", "postgres://admin:admin@localhost/user_service?sslmode=disable")
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
		}
	})
	return db
}
