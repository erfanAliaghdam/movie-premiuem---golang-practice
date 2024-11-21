package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword generates a bcrypt hash of the password.
func HashPassword(password string) (string, error) {
	// Generate a bcrypt hash with a cost of bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ComparePassword checks if the provided password matches the hashed password.
func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
