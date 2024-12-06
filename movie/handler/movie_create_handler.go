package handler

import (
	"log"
	"movie_premiuem/core"
	"movie_premiuem/core/utils"
	"movie_premiuem/movie/repository"
	"movie_premiuem/movie/service"
	"movie_premiuem/movie/validator"
	"net/http"
)

func MovieCreateHandler(w http.ResponseWriter, r *http.Request) {

	// check request method
	if r.Method != "POST" {
		utils.InvalidRequestMethod405(w)
		return
	}
	// validate body
	movieValidator := validator.NewMovieCreateValidator(r)
	isValid, errorFields := movieValidator.Validate()
	if !isValid {
		utils.BadRequestError400(w, "Bad Request.", errorFields)
		return
	}

	// get configs
	db := core.AppInstance.GetDB()

	// TODO check if user is admin ...

	movieRepository := repository.NewMovieRepository(db)
	movieService := service.NewMovieService(movieRepository)

	createdMovie, movieCreationErr := movieService.CreateMovie(
		movieValidator.Title,
		movieValidator.Description,
		&movieValidator.ImageFile,
	)
	if movieCreationErr != nil {
		log.Println("error occurred in MovieCreateHandler:", movieCreationErr)
		utils.InternalServerError500(w)
		return
	}

	response := utils.Response{
		Status:  "success",
		Message: "Movie created successfully.",
		Code:    "",
		Data:    createdMovie,
	}
	utils.WriteJSONResponse(w, response, http.StatusCreated)
}
