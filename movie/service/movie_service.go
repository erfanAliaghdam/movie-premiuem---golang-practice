package service

import (
	"bytes"
	"movie_premiuem/core/utils"
	"movie_premiuem/movie/entity"
	"movie_premiuem/movie/repository"
	"time"
)

type MovieService interface {
	CreateMovie(title string, description string, imageFile *[]byte) (entity.Movie, error)
}

type movieService struct {
	movieRepository repository.MovieRepository
}

func NewMovieService(movieRepo repository.MovieRepository) MovieService {
	if movieRepo == nil {
		panic("movieRepo cannot be nil")
	}
	return &movieService{movieRepository: movieRepo}
}

func (s *movieService) CreateMovie(title string, description string, imageFile *[]byte) (entity.Movie, error) {
	fileName := string(time.Now().Unix()) + title
	url, uploadErr := utils.UploadFileToBucket(bytes.NewReader(*imageFile), fileName)
	if uploadErr != nil {
		return entity.Movie{}, uploadErr
	}

	createdMovie, movieCreationErr := s.movieRepository.CreateMovie(title, description, url)
	if movieCreationErr != nil {
		return entity.Movie{}, movieCreationErr
	}

	return createdMovie, nil
}
