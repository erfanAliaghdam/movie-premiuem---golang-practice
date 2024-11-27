package app

import (
	"database/sql"
	"log"
	"sync"
)

// Application interface
type Application interface {
	GetDB() *sql.DB
	CloseDB()
}

// application struct implementing Application
type application struct {
	db *sql.DB
}

// Ensure thread-safe initialization with sync.Once
var (
	AppInstance Application
	once        sync.Once
)

// NewApplication initializes the Application with a database connection
func NewApplication(dbName string) Application {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	return &application{db: db}
}

// InitApplication initializes the global Application instance
func InitApplication(dbName string) {
	once.Do(func() {
		AppInstance = NewApplication(dbName)
	})
}

// GetDB method to retrieve the database connection
func (a *application) GetDB() *sql.DB {
	return a.db
}

func (a *application) CloseDB() {
	if a.db != nil {
		a.db.Close()
	}
}
