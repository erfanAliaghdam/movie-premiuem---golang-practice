package repositories

import (
	"database/sql"
	"movie_premiuem/utils"
)

type UserRepository interface {
	CreateUser(Email string, Password string) (int64, error)
	CheckIfUserExistsByEmail(Email string) (bool, error)
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
func (r *userRepository) CreateUser(Email string, Password string) (int64, error) {
	hashedPassword, hashError := r.hash.HashPassword(Password)
	if hashError != nil {
		return 0, hashError
	}

	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	result, err := r.db.Exec(query, Email, hashedPassword)
	if err != nil {
		return 0, err
	}

	// Get the last inserted ID to update the user entity
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *userRepository) CheckIfUserExistsByEmail(Email string) (bool, error) {
	query := "SELECT id FROM users WHERE email = ?"
	row := r.db.QueryRow(query, Email)

	var id int64
	err := row.Scan(&id)
	if err != nil {
		return false, err
	}

	return true, nil
}
