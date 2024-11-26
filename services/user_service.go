package services

import (
	"fmt"
	"movie_premiuem/entity"
	"movie_premiuem/entity/repositories"
)

type UserService interface {
	RegisterUser(user entity.User) (entity.User, error)
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

func (s *userService) RegisterUser(user entity.User) (entity.User, error) {
	// Save the user to the database
	cratedUser, userCreationErr := s.userRepository.CreateUser(user)
	if userCreationErr != nil {
		return entity.User{}, fmt.Errorf("failed to register user: %w", userCreationErr)
	}

	return cratedUser, nil
}
