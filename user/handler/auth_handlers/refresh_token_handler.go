package auth_handlers

import (
	"encoding/json"
	"log"
	"movie_premiuem/core/utils"
	"net/http"
)

func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.InvalidRequestMethod405(w)
		return
	}

	// Parse the request to get the refresh token
	var requestData struct {
		RefreshToken string `json:"refresh_token"`
	}

	// Parse the request body
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestData)
	if err != nil || requestData.RefreshToken == "" {
		utils.BadRequestError400(w, "Refresh token is required.", nil)
		return
	}

	// Validate the refresh token and generate a new access token
	userID, tokenErr := utils.VerifyRefreshToken(requestData.RefreshToken)
	if tokenErr != nil {
		utils.UnauthorizedError401(w)
		return
	}

	accessToken, refreshToken, tokenErr := utils.GenerateJWT(userID)
	if tokenErr != nil {
		log.Println(tokenErr)
		utils.InternalServerError500(w)
		return
	}

	response := utils.Response{
		Status:  "success",
		Message: "Token refreshed successfully.",
		Code:    "",
		Data: map[string]string{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		},
	}

	utils.WriteJSONResponse(w, response, http.StatusOK)
}
