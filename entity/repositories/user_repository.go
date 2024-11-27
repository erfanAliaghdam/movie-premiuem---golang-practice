package repositories

import (
	"database/sql"
	"movie_premiuem/entity"
	"movie_premiuem/utils"
)

type UserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
}

type userRepository struct {
	db   *sql.DB
	hash utils.HashFactory
}

// NewUserRepository creates a new instance of userRepository
func NewUserRepository(db *sql.DB, hash ...utils.HashFactory) UserRepository {
	var hashFac utils.HashFactory
	if len(hash) == 0 {
		hashFac = utils.NewHashFactory(25)
	} else {
		hashFac = hash[0]
	}
	return &userRepository{db: db, hash: hashFac}
}

// CreateUser inserts a new user into the database
func (r *userRepository) CreateUser(user entity.User) (entity.User, error) {
	hashedPassword, hashError := r.hash.HashPassword(user.Password)
	if hashError != nil {
		return entity.User{}, hashError
	}

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
	user.Password = hashedPassword

	return user, nil
}
