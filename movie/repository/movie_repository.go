package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"log"
	"movie_premiuem/core"
	"movie_premiuem/core/custom_errors"
	"movie_premiuem/movie/entity"
	combinedEntity "movie_premiuem/movie/entity/combined_eneities"
)

//go:generate mockery --name=MovieRepository
type MovieRepository interface {
	GetMovieList() ([]combinedEntity.MovieWithUrls, error)
	GetMovieDetail(movieID int64) (combinedEntity.MovieWithUrls, error)
	CreateMovie(title string, description string, imageUrl string) (entity.Movie, error)
	SeedData() error
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

func (r *movieRepository) CreateMovie(title string, description string, imageUrl string) (entity.Movie, error) {
	query := `INSERT INTO movies (title, description, image_url) VALUES (?, ?, ?)
				RETURNING id, title, description, image_url, created_at;`

	var createdMovie entity.Movie

	err := r.db.QueryRow(
		query,
		title,
		description,
		imageUrl,
	).Scan(
		&createdMovie.ID,
		&createdMovie.Title,
		&createdMovie.Description,
		&createdMovie.ImageUrl,
		&createdMovie.CreatedAt,
	)
	if err != nil {
		return entity.Movie{}, err
	}

	return createdMovie, nil

}

func (r *movieRepository) GetMovieDetail(movieID int64) (combinedEntity.MovieWithUrls, error) {
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
		WHERE 
			movies.id = ?
		GROUP BY 
			movies.id;
	`

	var movie combinedEntity.MovieWithUrls
	var urlsJSON string

	row := r.db.QueryRow(query, movieID)
	// Scan the result into the movie struct
	scanErr := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.CreatedAt,
		&movie.Description,
		&movie.ImageUrl,
		&urlsJSON,
	)
	if scanErr != nil {
		if errors.Is(scanErr, sql.ErrNoRows) {
			// not found
			return combinedEntity.MovieWithUrls{}, custom_errors.NotExists
		}
		return combinedEntity.MovieWithUrls{}, scanErr
	}

	err := json.Unmarshal([]byte(urlsJSON), &movie.Urls)
	if err != nil {
		return combinedEntity.MovieWithUrls{}, errors.New("failed to parse URLs JSON: " + err.Error())
	}
	if movie.Urls[0].ID == 0 {
		movie.Urls = []entity.MovieUrl{}
	}

	return movie, nil
}

func (r *movieRepository) SeedData() error {
	db := core.AppInstance.GetDB()
	// Generate sample data
	for i := 0; i < 5; i++ {
		title := gofakeit.Sentence(3)
		description := gofakeit.Paragraph(1, 3, 5, "")
		imageURL := gofakeit.URL()

		// Insert movie
		result, err := db.Exec("INSERT INTO movies (title, description, image_url) VALUES (?, ?, ?)",
			title, description, imageURL)
		if err != nil {
			log.Fatal(err)
		}

		// Get movie ID
		movieID, _ := result.LastInsertId()

		// Insert related URLs
		for j := 0; j < 3; j++ {
			urlTitle := gofakeit.Sentence(2)
			url := gofakeit.URL()

			_, err = db.Exec("INSERT INTO movie_urls (title, url, movie_id) VALUES (?, ?, ?)",
				urlTitle, url, movieID)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	fmt.Println("Random data inserted using GoFakeIt!")
	return nil
}
