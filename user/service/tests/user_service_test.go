package tests

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"movie_premiuem/core/custom_errors"
	"movie_premiuem/user/repository/mocks"
	"movie_premiuem/user/service"
	"testing"
)

func TestUserService_RegisterUser_FailOnExistingUser(t *testing.T) {
	testUserMail := "example@example.com"
	testUserPass := "password"
	userRepositoryMock := &mocks.UserRepository{}

	userRepositoryMock.On(
		"CheckIfUserExistsByEmail",
		testUserMail,
	).Return(true, nil)

	userService := service.NewUserService(userRepositoryMock)
	_, err := userService.RegisterUser(testUserMail, testUserPass)

	assert.Error(t, err)
	assert.Equal(t, err.Error(), custom_errors.AlreadyExists.Error())

	userRepositoryMock.AssertExpectations(t)
}

func TestUserService_RegisterUser_FailOnCheckingUser(t *testing.T) {
	testUserMail := "example@example.com"
	testUserPass := "password"
	userRepositoryMock := &mocks.UserRepository{}

	userRepositoryMock.On(
		"CheckIfUserExistsByEmail",
		testUserMail,
	).Return(false, errors.New("error"))

	userService := service.NewUserService(userRepositoryMock)
	_, err := userService.RegisterUser(testUserMail, testUserPass)

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "error")

	userRepositoryMock.AssertExpectations(t)
}

func TestUserService_RegisterUser_FailOnCreateUserQuery(t *testing.T) {
	testUserMail := "example@example.com"
	testUserPass := "password"
	userRepositoryMock := &mocks.UserRepository{}

	userRepositoryMock.On(
		"CheckIfUserExistsByEmail",
		testUserMail,
	).Return(false, nil).Once()

	userRepositoryMock.On(
		"CreateUser",
		testUserMail,
		testUserPass,
	).Return(int64(0), errors.New("error")).Once()

	userService := service.NewUserService(userRepositoryMock)
	_, err := userService.RegisterUser(testUserMail, testUserPass)

	assert.Error(t, err)
	userRepositoryMock.AssertExpectations(t)
}

func TestUserService_RegisterUser_Success(t *testing.T) {
	testUserMail := "example@example.com"
	testUserPass := "password"
	userRepositoryMock := &mocks.UserRepository{}

	userRepositoryMock.On(
		"CheckIfUserExistsByEmail",
		testUserMail,
	).Return(false, nil).Once()

	userRepositoryMock.On(
		"CreateUser",
		testUserMail,
		testUserPass,
	).Return(int64(1), nil).Once()

	userService := service.NewUserService(userRepositoryMock)
	userID, err := userService.RegisterUser(testUserMail, testUserPass)

	assert.Equal(t, int64(1), userID)
	assert.NoError(t, err)

	userRepositoryMock.AssertExpectations(t)
}
