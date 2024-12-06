package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"movie_premiuem/movie/entity"
	combinedEntity "movie_premiuem/movie/entity/combined_eneities"
)

type MovieRepository interface {
	GetMovieList() ([]combinedEntity.MovieWithUrls, error)
}

type movieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) MovieRepository {
	return &movieRepository{db: db}
}

func (r *movieRepository) GetMovieList() ([]combinedEntity.MovieWithUrls, error) {
	query := `
		SELECT 
			movies.id AS id,
			movies.title AS title,
			movies.created_at AS created_at,
			movies.description AS description,
			movies.image_url AS image_url,
			COALESCE(
				json_group_array(
					json_object(
						'id', movie_urls.id,
						'url', movie_urls.url,
						'title', movie_urls.title,
						'created_at', strftime('%Y-%m-%dT%H:%M:%SZ', movie_urls.created_at)
					)
				), 
				'[]'
			) AS urls
		FROM 
			movies
		LEFT JOIN 
			movie_urls 
		ON 
			movies.id = movie_urls.movie_id
		GROUP BY 
			movies.id;
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []combinedEntity.MovieWithUrls

	for rows.Next() {
		var movie combinedEntity.MovieWithUrls
		var urlsJSON string

		// Scan the result into variables
		scanErr := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.CreatedAt,
			&movie.Description,
			&movie.ImageUrl,
			&urlsJSON,
		)
		if scanErr != nil {
			return nil, scanErr
		}

		err = json.Unmarshal([]byte(urlsJSON), &movie.Urls)
		if err != nil {
			return nil, errors.New("failed to parse URLs JSON: " + err.Error())
		}
		if movie.Urls[0].ID == 0 {
			movie.Urls = []entity.MovieUrl{}
		}

		movies = append(movies, movie)
	}

	// Check for errors during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}
