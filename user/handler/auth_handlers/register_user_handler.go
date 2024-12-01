package auth_handlers

import (
	"errors"
	"log"
	"movie_premiuem/core"
	"movie_premiuem/core/custom_errors"
	"movie_premiuem/core/utils"
	"movie_premiuem/user/repository"
	"movie_premiuem/user/serializer/auth_serializers"
	"movie_premiuem/user/service"
	"net/http"
)

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	db := core.AppInstance.GetDB()

	if r.Method != "POST" {
		utils.InvalidRequestMethod405(w)
	}

	serializer := auth_serializers.NewRegisterUserSerializer(r)
	isValid, errorFields := serializer.Serialize()
	if !isValid {
		utils.BadRequestError400(w, "Bad request.", errorFields)
		return
	}

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	_, registerUserErr := userService.RegisterUser(serializer.Email, serializer.Password)
	if registerUserErr != nil {
		if errors.Is(registerUserErr, custom_errors.AlreadyExists) {
			utils.BadRequestError400(w, "User already exists.", nil)
			return
		}
		log.Println(registerUserErr)
		utils.InternalServerError500(w)
		return
	}

	response := utils.Response{
		Status:  "success",
		Message: "user registered successfully.",
		Code:    "",
		Data:    nil,
	}

	utils.WriteJSONResponse(w, response, http.StatusCreated)
}
