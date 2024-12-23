package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Code    string      `json:"code"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Code    string      `json:"code"`
	Errors  interface{} `json:"errors"`
}

func WriteJSONResponse(w http.ResponseWriter, response Response, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// Fixed Responses

func InternalServerError500(w http.ResponseWriter) {
	WriteJSONResponse(w, Response{
		Status:  "fail",
		Message: "An unexpected error occurred. Please try again later.",
		Code:    "INTERNAL_SERVER_ERROR",
		Data:    nil,
	}, http.StatusInternalServerError)
}

func UnauthorizedError401(w http.ResponseWriter) {
	WriteJSONResponse(w, Response{
		Status:  "fail",
		Message: "Unauthorized access. Please provide valid credentials.",
		Code:    "UNAUTHORIZED",
		Data:    nil,
	}, http.StatusUnauthorized)
}

func ForbiddenError403(w http.ResponseWriter) {
	WriteJSONResponse(w, Response{
		Status:  "fail",
		Message: "You do not have permission to access this resource.",
		Code:    "FORBIDDEN",
		Data:    nil,
	}, http.StatusForbidden)
}

func BadRequestError400(w http.ResponseWriter, message string, fieldErrors interface{}) {
	response := map[string]interface{}{
		"status":  "failed",
		"message": message,
	}
	if fieldErrors != nil {
		response["errors"] = fieldErrors
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(response)
}

func NotFoundError404(w http.ResponseWriter) {
	WriteJSONResponse(w, Response{
		Status:  "fail",
		Message: "The requested resource was not found.",
		Code:    "NOT_FOUND",
		Data:    nil,
	}, http.StatusNotFound)
}

func InvalidRequestMethod405(w http.ResponseWriter) {
	WriteJSONResponse(w, Response{
		Status:  "fail",
		Message: "Invalid request method.",
		Code:    "",
		Data:    nil,
	}, http.StatusMethodNotAllowed)
}
