package auth_serializers

import (
	"encoding/json"
	"movie_premiuem/core/utils"
	"net/http"
)

type LoginUserValidator struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (v *LoginUserValidator) Validate() (bool, map[string]string) {
	return utils.ValidateField(v)
}

// NewLoginUserValidator NewRegisterSerializer Method to deserialize the request body into a RegisterValidator
func NewLoginUserValidator(r *http.Request) *LoginUserValidator {
	var validator LoginUserValidator
	json.NewDecoder(r.Body).Decode(&validator)

	return &validator
}
