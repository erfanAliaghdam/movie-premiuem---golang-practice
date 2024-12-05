package core

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	SecretKey string
	DBName    string
	RedisAddr string
}

// LoadConfig loads environment variables into a Config struct
func LoadConfig() Config {
	// Load .env file for development
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	return Config{
		SecretKey: getEnv("SECRET_KEY", "default-secret-key"),
		DBName:    getEnv("DB_NAME", "db"),
		RedisAddr: getEnv("REDIS_URL", "localhost:6379"),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
