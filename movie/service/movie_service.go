package service

import (
	"bytes"
	"movie_premiuem/core/utils"
	"movie_premiuem/movie/entity"
	"movie_premiuem/movie/repository"
	"strconv"
	"time"
)

type MovieService interface {
	SetBucket(bucket utils.Bucket)
	CreateMovie(title string, description string, imageFile *[]byte) (entity.Movie, error)
}

type movieService struct {
	movieRepository repository.MovieRepository
	bucket          utils.Bucket
}

func NewMovieService(movieRepo repository.MovieRepository) MovieService {
	if movieRepo == nil {
		panic("movieRepo cannot be nil")
	}
	return &movieService{movieRepository: movieRepo, bucket: &utils.S3Bucket{}}
}

func (s *movieService) SetBucket(bucket utils.Bucket) {
	s.bucket = bucket
}

func (s *movieService) CreateMovie(title string, description string, imageFile *[]byte) (entity.Movie, error) {
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + title

	url, uploadErr := s.bucket.UploadFileToBucket(bytes.NewReader(*imageFile), fileName)
	if uploadErr != nil {
		return entity.Movie{}, uploadErr
	}

	createdMovie, movieCreationErr := s.movieRepository.CreateMovie(title, description, url)
	if movieCreationErr != nil {
		return entity.Movie{}, movieCreationErr
	}

	return createdMovie, nil
}
