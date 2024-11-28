package services

import (
	"fmt"
	"movie_premiuem/custom_errors"
	"movie_premiuem/entity/repositories"
)

type UserService interface {
	RegisterUser(email string, password string) (int64, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

// NewUserService initializes a new UserService
func NewUserService(userRepo repositories.UserRepository) UserService {
	if userRepo == nil {
		panic("userRepo cannot be nil")
	}
	return &userService{userRepository: userRepo}
}

func (s *userService) RegisterUser(email string, password string) (int64, error) {
	// check if email exists or not
	exists, userCheckError := s.userRepository.CheckIfUserExistsByEmail(email)
	if userCheckError != nil {
		return 0, userCheckError
	}
	if exists {
		return 0, custom_errors.AlreadyExists
	}
	// Save the user to the database
	cratedUserID, userCreationErr := s.userRepository.CreateUser(email, password)
	if userCreationErr != nil {
		return 0, fmt.Errorf("failed to register user: %w", userCreationErr)
	}

	return cratedUserID, nil
}
