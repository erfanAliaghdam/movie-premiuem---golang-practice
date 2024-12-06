package auth_serializers

import (
	"encoding/json"
	"movie_premiuem/core/utils"
	"net/http"
)

type RegisterUserValidator struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (v *RegisterUserValidator) Validate() (bool, map[string]string) {
	// Validate fields
	return utils.ValidateField(v)
}

// NewRegisterUserValidator NewRegisterSerializer Method to deserialize the request body into a RegisterSerializer
func NewRegisterUserValidator(r *http.Request) *RegisterUserValidator {
	var validator RegisterUserValidator
	json.NewDecoder(r.Body).Decode(&validator)

	return &validator
}
