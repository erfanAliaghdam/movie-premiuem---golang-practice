package utils

import "golang.org/x/crypto/bcrypt"

type HashFactory interface {
	HashPassword(password string) (string, error)
	CompareHashAndPassword(hashedPassword string, password string) bool
}

type hashFactory struct {
	cost int
}

func NewHashFactory(cost int) HashFactory {
	return &hashFactory{cost: cost}
}

// HashPassword generates a bcrypt hash of the password.
func (h *hashFactory) HashPassword(password string) (string, error) {
	// Generate a bcrypt hash with a cost of bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CompareHashAndPassword checks if the provided password matches the hashed password.
func (h *hashFactory) CompareHashAndPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
