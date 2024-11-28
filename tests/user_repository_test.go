package tests

import (
	"errors"
	"movie_premiuem/tests/mocks/mock_utils"
	"movie_premiuem/user/entity"
	"movie_premiuem/user/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryCreateUser(t *testing.T) {
	// Mock the database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Create a mock HashFactory
	mockHash := new(mock_utils.MockHashFactory)
	mockHash.On(
		"HashPassword",
		"plain_password",
	).Return(
		"hashed_password",
		nil,
	)

	// Create a repository instance
	repo := repository.NewUserRepository(db, mockHash)

	// Define the user entity
	inputUser := entity.User{
		Email:    "test@example.com",
		Password: "plain_password",
	}
	expectedUser := entity.User{
		ID:       1,
		Email:    "test@example.com",
		Password: "hashed_password",
	}

	t.Run("success - user created", func(t *testing.T) {
		// Expect query to be executed
		mock.ExpectExec("INSERT INTO users").
			WithArgs(inputUser.Email, "plain_password").
			WillReturnResult(sqlmock.NewResult(1, 1)) // Mock the last insert ID as 1

		// Call the CreateUser method
		createdUser, err := repo.CreateUser(inputUser)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, expectedUser, createdUser)
		mockHash.AssertExpectations(t)
	})

	t.Run("failure - db error", func(t *testing.T) {
		// Mock database error
		mock.ExpectExec("INSERT INTO users").
			WithArgs(inputUser.Email, "plain_password").
			WillReturnError(errors.New("db error"))

		// Call the CreateUser method
		_, err := repo.CreateUser(inputUser)

		// Assertions
		assert.Error(t, err)
		assert.EqualError(t, err, "db error")
		mockHash.AssertExpectations(t)
	})
}
