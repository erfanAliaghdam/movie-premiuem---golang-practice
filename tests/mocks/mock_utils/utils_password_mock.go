package mock_utils

import "github.com/stretchr/testify/mock"

type MockHashFactory struct {
	mock.Mock
}

// HashPassword mocks the HashPassword method
func (m *MockHashFactory) HashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

// CompareHashAndPassword mocks the CompareHashAndPassword method
func (m *MockHashFactory) CompareHashAndPassword(hashedPassword string, password string) bool {
	args := m.Called(hashedPassword, password)
	return args.Bool(0)
}
