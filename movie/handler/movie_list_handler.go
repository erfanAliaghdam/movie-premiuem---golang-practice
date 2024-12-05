package handler

import (
	"log"
	"movie_premiuem/core"
	"movie_premiuem/core/utils"
	"movie_premiuem/movie/repository"
	userRepo "movie_premiuem/user/repository"
	"net/http"
	"strings"
)

func MovieListHandler(w http.ResponseWriter, r *http.Request) {

	// check request method
	if r.Method != "GET" {
		utils.InvalidRequestMethod405(w)
		return
	}

	// get configs
	db := core.AppInstance.GetDB()

	// check if user is authenticated
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		log.Println("Authorization header is missing")
		utils.UnauthorizedError401(w)
		return
	}

	token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
	if token == "" {
		log.Println("Bearer token is missing in the Authorization header")
		utils.UnauthorizedError401(w)
		return
	}

	userID, tokenErr := utils.VerifyToken(token)
	if tokenErr != nil {
		utils.UnauthorizedError401(w)
		return
	}

	// check permissions
	userLicenseRepository := userRepo.NewUserLicenseRepository(db)
	userHasLicense, userLicenseErr := userLicenseRepository.CheckIfUserHasActiveLicense(userID)
	if userLicenseErr != nil {
		log.Println(userLicenseErr)
		utils.InternalServerError500(w)
		return
	}
	if !userHasLicense {
		utils.ForbiddenError403(w)
		return
	}

	movieRepository := repository.NewMovieRepository(db)
	movieList, err := movieRepository.GetMovieList()
	if err != nil {
		log.Println("error occurred on MovieListHandler:", err)
		utils.InternalServerError500(w)
		return
	}

	response := utils.Response{
		Status:  "success",
		Message: "movie list fetched successfully",
		Data:    movieList,
		Code:    "",
	}
	utils.WriteJSONResponse(w, response, http.StatusOK)
}
