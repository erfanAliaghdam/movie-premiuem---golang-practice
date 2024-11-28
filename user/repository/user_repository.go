package repository

import (
	"database/sql"
	"movie_premiuem/core/custom_errors"
	"movie_premiuem/core/utils"
	"movie_premiuem/user/entity"
)

type UserRepository interface {
	CreateUser(Email string, Password string) (int64, error)
	CheckIfUserExistsByEmail(Email string) (bool, error)
	GetUserByEmail(Email string) (entity.User, error)
	ValidateUserByEmailAndPassword(Email string, Password string) (bool, int64, error)
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
		if err == sql.ErrNoRows {
			// No rows were found; user does not exist
			return false, nil
		}
		// An actual error occurred
		return false, err
	}

	// User exists
	return true, nil
}

func (r *userRepository) GetUserByEmail(Email string) (entity.User, error) {
	query := "SELECT * FROM users WHERE email = ?"

	row := r.db.QueryRow(query, Email)
	var user entity.User
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, custom_errors.NotExists
		}
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) ValidateUserByEmailAndPassword(Email string, Password string) (bool, int64, error) {
	user, getUserErr := r.GetUserByEmail(Email)
	if getUserErr != nil {
		return false, 0, getUserErr
	}

	return r.hash.CompareHashAndPassword(user.Password, Password), user.ID, nil
}
