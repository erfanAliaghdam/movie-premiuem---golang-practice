package auth_handlers

import (
	"log"
	"movie_premiuem/core"
	"movie_premiuem/core/utils"
	"movie_premiuem/user/repository"
	"movie_premiuem/user/validator/auth_serializers"
	"net/http"
)

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	db := core.AppInstance.GetDB()

	if r.Method != "POST" {
		utils.InvalidRequestMethod405(w)
		return
	}

	validator := auth_serializers.NewLoginUserValidator(r)
	isValid, errorFields := validator.Validate()
	if !isValid {
		utils.BadRequestError400(w, "Bad Request.", errorFields)
		return
	}

	userRepository := repository.NewUserRepository(db)
	userIsValid, userId, userValidationErr := userRepository.ValidateUserByEmailAndPassword(
		validator.Email,
		validator.Password,
	)
	if userValidationErr != nil || !userIsValid {
		log.Println("error occurred in user login handler :", userValidationErr)
		utils.UnauthorizedError401(w)
		return
	}

	access, refresh, tokenErr := utils.GenerateJWT(userId)
	if tokenErr != nil {
		log.Print("error occurred in user login handler :", tokenErr)
		utils.InternalServerError500(w)
		return
	}

	response := utils.Response{
		Status:  "success",
		Message: "user logged in successfully",
		Data: map[string]string{
			"access":  access,
			"refresh": refresh,
		},
	}
	utils.WriteJSONResponse(w, response, http.StatusOK)
}
