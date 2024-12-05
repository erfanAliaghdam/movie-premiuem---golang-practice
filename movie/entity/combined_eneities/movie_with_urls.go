package entity

import "movie_premiuem/movie/entity"

type MovieWithUrls struct {
	entity.Movie
	Urls []entity.MovieUrl `json:"urls"`
}
