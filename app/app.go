package app

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3" // for SQLite support
	"github.com/redis/go-redis/v9"
)

// Application interface
type Application interface {
	GetDB() *sql.DB
	GetRedis() *redis.Client
	CloseDB()
	CloseRedis()
}

// application struct implementing Application
type application struct {
	db          *sql.DB
	redisClient *redis.Client
}

// Ensure thread-safe initialization with sync.Once
var (
	AppInstance Application
	once        sync.Once
)

// NewApplication initializes the Application with a database and Redis connection
func NewApplication(dbName string, redisAddr string) Application {
	// Initialize Database
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Initialize Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisAddr, // e.g., "localhost:6379"
	})

	return &application{
		db:          db,
		redisClient: redisClient,
	}
}

// InitApplication initializes the global Application instance
func InitApplication(dbName string, redisAddr string) {
	once.Do(func() {
		AppInstance = NewApplication(dbName, redisAddr)
	})
}

// GetDB retrieves the database connection
func (a *application) GetDB() *sql.DB {
	return a.db
}

// GetRedis retrieves the Redis client
func (a *application) GetRedis() *redis.Client {
	return a.redisClient
}

// CloseDB closes the database connection
func (a *application) CloseDB() {
	if a.db != nil {
		a.db.Close()
	}
}

// CloseRedis closes the Redis connection
func (a *application) CloseRedis() {
	if a.redisClient != nil {
		a.redisClient.Close()
	}
}
