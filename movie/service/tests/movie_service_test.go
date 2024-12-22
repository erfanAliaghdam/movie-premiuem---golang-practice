package tests

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	utilsMocks "movie_premiuem/core/utils/mocks"
	"movie_premiuem/movie/entity"
	movieRepositoryMock "movie_premiuem/movie/repository/mocks"
	"movie_premiuem/movie/service"
	"testing"
	"time"
)

func TestUserService_CreateMovie_FailBucket(t *testing.T) {
	// Set up test data
	testMovieTitle := "test movie title"
	testMovieDescription := "test movie description"
	testByteFileContent := []byte("test content")
	testFileContent := bytes.NewReader(testByteFileContent)

	// Set up the mocks
	movieRepository := &movieRepositoryMock.MovieRepository{}
	bucketMock := &utilsMocks.Bucket{}
	bucketMock.On(
		"UploadFileToBucket",
		testFileContent,
		mock.Anything,
	).Return(
		"",
		errors.New("bucket failed"),
	).Once()

	// Create service with mock dependencies
	movieServiceOBJ := service.NewMovieService(movieRepository)
	movieServiceOBJ.SetBucket(bucketMock)

	// Call the method and capture the error
	_, err := movieServiceOBJ.CreateMovie(testMovieTitle, testMovieDescription, &testByteFileContent)

	// Assert that an error occurred (assuming the service will return an error if bucket fails)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "bucket failed")

	// Assert expectations on the mock
	bucketMock.AssertExpectations(t)
}

func TestUserService_CreateMovie_FailRepository(t *testing.T) {
	// Set up test data
	testMovieTitle := "test movie title"
	testMovieDescription := "test movie description"
	testByteFileContent := []byte("test content")
	testMovieS3Url := "test url"
	testFileContent := bytes.NewReader(testByteFileContent)

	// Set up the mocks
	movieRepository := &movieRepositoryMock.MovieRepository{}
	bucketMock := &utilsMocks.Bucket{}
	bucketMock.On(
		"UploadFileToBucket",
		testFileContent,
		mock.Anything,
	).Return(
		testMovieS3Url,
		nil,
	).Once()

	movieRepository.On(
		"CreateMovie",
		testMovieTitle,
		testMovieDescription,
		testMovieS3Url,
	).Return(
		entity.Movie{},
		errors.New("repository failed"),
	).Once()

	// Create service with mock dependencies
	movieServiceOBJ := service.NewMovieService(movieRepository)
	movieServiceOBJ.SetBucket(bucketMock)

	// Call the method and capture the error
	_, err := movieServiceOBJ.CreateMovie(testMovieTitle, testMovieDescription, &testByteFileContent)

	// Assert that an error occurred (assuming the service will return an error if bucket fails)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "repository failed")

	// Assert expectations on the mock
	bucketMock.AssertExpectations(t)
	movieRepository.AssertExpectations(t)
}

func TestUserService_CreateMovie_Success(t *testing.T) {
	// Set up test data
	testMovieTitle := "test movie title"
	testMovieDescription := "test movie description"
	testByteFileContent := []byte("test content")
	testMovieS3Url := "test url"
	testFileContent := bytes.NewReader(testByteFileContent)

	// Set up the mocks
	movieRepository := &movieRepositoryMock.MovieRepository{}
	bucketMock := &utilsMocks.Bucket{}
	bucketMock.On(
		"UploadFileToBucket",
		testFileContent,
		mock.Anything,
	).Return(
		testMovieS3Url,
		nil,
	).Once()

	movieRepository.On(
		"CreateMovie",
		testMovieTitle,
		testMovieDescription,
		testMovieS3Url,
	).Return(
		entity.Movie{
			ID:          1,
			Title:       testMovieTitle,
			Description: testMovieDescription,
			CreatedAt:   time.Now(),
			ImageUrl:    testMovieS3Url,
		},
		nil,
	).Once()

	// Create service with mock dependencies
	movieServiceOBJ := service.NewMovieService(movieRepository)
	movieServiceOBJ.SetBucket(bucketMock)

	// Call the method and capture the error
	movie, err := movieServiceOBJ.CreateMovie(testMovieTitle, testMovieDescription, &testByteFileContent)

	// Assert no error and the movie was created correctly
	assert.NoError(t, err)
	assert.Equal(t, movie.Title, testMovieTitle)
	assert.Equal(t, movie.Description, testMovieDescription)
	assert.Equal(t, movie.ImageUrl, testMovieS3Url)

	// Assert expectations on the mock
	bucketMock.AssertExpectations(t)
	movieRepository.AssertExpectations(t)
}
