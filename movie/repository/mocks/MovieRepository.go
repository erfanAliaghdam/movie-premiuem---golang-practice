// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	entity "movie_premiuem/movie/entity"
	combined_eneities "movie_premiuem/movie/entity/combined_eneities"

	mock "github.com/stretchr/testify/mock"
)

// MovieRepository is an autogenerated mock type for the MovieRepository type
type MovieRepository struct {
	mock.Mock
}

// CreateMovie provides a mock function with given fields: title, description, imageUrl
func (_m *MovieRepository) CreateMovie(title string, description string, imageUrl string) (entity.Movie, error) {
	ret := _m.Called(title, description, imageUrl)

	if len(ret) == 0 {
		panic("no return value specified for CreateMovie")
	}

	var r0 entity.Movie
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (entity.Movie, error)); ok {
		return rf(title, description, imageUrl)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) entity.Movie); ok {
		r0 = rf(title, description, imageUrl)
	} else {
		r0 = ret.Get(0).(entity.Movie)
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(title, description, imageUrl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMovieDetail provides a mock function with given fields: movieID
func (_m *MovieRepository) GetMovieDetail(movieID int64) (combined_eneities.MovieWithUrls, error) {
	ret := _m.Called(movieID)

	if len(ret) == 0 {
		panic("no return value specified for GetMovieDetail")
	}

	var r0 combined_eneities.MovieWithUrls
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (combined_eneities.MovieWithUrls, error)); ok {
		return rf(movieID)
	}
	if rf, ok := ret.Get(0).(func(int64) combined_eneities.MovieWithUrls); ok {
		r0 = rf(movieID)
	} else {
		r0 = ret.Get(0).(combined_eneities.MovieWithUrls)
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(movieID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMovieList provides a mock function with no fields
func (_m *MovieRepository) GetMovieList() ([]combined_eneities.MovieWithUrls, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetMovieList")
	}

	var r0 []combined_eneities.MovieWithUrls
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]combined_eneities.MovieWithUrls, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []combined_eneities.MovieWithUrls); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]combined_eneities.MovieWithUrls)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SeedData provides a mock function with no fields
func (_m *MovieRepository) SeedData() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for SeedData")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMovieRepository creates a new instance of MovieRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMovieRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MovieRepository {
	mock := &MovieRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
