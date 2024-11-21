package repositories

import (
	"database/sql"
	"movie_premiuem/entity"
)

type UserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
}

type userRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new instance of userRepository
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

// CreateUser inserts a new user into the database
func (r *userRepository) CreateUser(user entity.User) (entity.User, error) {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	result, err := r.db.Exec(query, user.Email, user.Password)
	if err != nil {
		return entity.User{}, err
	}

	// Get the last inserted ID to update the user entity
	id, err := result.LastInsertId()
	if err != nil {
		return entity.User{}, err
	}
	user.ID = id

	return user, nil
}
