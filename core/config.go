package core

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	SecretKey       string
	DBName          string
	RedisAddr       string
	BucketAccessKey string
	BucketSecretKey string
	BucketEndpoint  string
	BucketName      string
}

// LoadConfig loads environment variables into a Config struct
func LoadConfig() Config {
	// Load .env file for development
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}
	SecretKey, _ := getEnv("SECRET_KEY", "default-secret-key", false)
	DBNAME, _ := getEnv("DB_NAME", "db", false)
	RedisAddress, _ := getEnv("REDIS_URL", "localhost:6379", false)
	BucketAccessKey, err := getEnv("BUCKET_ACCESS_KEY", "", true)
	BucketSecret, err := getEnv("BUCKET_SECRET_KEY", "", true)
	BucketEndpoint, err := getEnv("BUCKET_ENDPOINT", "", true)
	BucketName, err := getEnv("BUCKET_NAME", "", true)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return Config{
		SecretKey:       SecretKey,
		DBName:          DBNAME,
		RedisAddr:       RedisAddress,
		BucketAccessKey: BucketAccessKey,
		BucketSecretKey: BucketSecret,
		BucketEndpoint:  BucketEndpoint,
		BucketName:      BucketName,
	}
}

func getEnv(key string, defaultValue string, errorOnNotFound bool) (string, error) {
	value, exists := os.LookupEnv(key)
	if errorOnNotFound && !exists {
		errorText := "error on config load from .env file:" + key
		return "", errors.New(errorText)
	}
	if exists {
		return value, nil
	}
	return defaultValue, nil
}
