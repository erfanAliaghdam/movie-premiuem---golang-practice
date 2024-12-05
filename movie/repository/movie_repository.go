package repository

import "movie_premiuem/movie/entity"

type MovieRepository interface {
	GetMovieList() entity.Movie
}

type movieRepository struct {
}

func NewMovieRepository() MovieRepository {
	return &movieRepository{}
}
