package handler

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"movie_premiuem/core"
	"movie_premiuem/core/custom_errors"
	"movie_premiuem/core/utils"
	"movie_premiuem/movie/repository"
	"net/http"
	"strconv"
)

func MovieDetailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		utils.InvalidRequestMethod405(w)
		return
	}

	movieIDParam := chi.URLParam(r, "id")
	movieID, movieIDErr := strconv.ParseInt(movieIDParam, 10, 64)
	if movieIDErr != nil {
		utils.NotFoundError404(w)
		return
	}

	db := core.AppInstance.GetDB()

	movieRepository := repository.NewMovieRepository(db)
	movie, movieDetailErr := movieRepository.GetMovieDetail(movieID)
	if movieDetailErr != nil {
		if errors.Is(movieDetailErr, custom_errors.NotExists) {
			utils.NotFoundError404(w)
			return
		}

		utils.InternalServerError500(w)
		return
	}

	response := utils.Response{
		Status:  "success",
		Message: "Movie Detail Fetched Successfully.",
		Data:    movie,
		Code:    "",
	}
	utils.WriteJSONResponse(w, response, 200)

}
