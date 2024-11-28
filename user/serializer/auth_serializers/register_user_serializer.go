package auth_serializers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"movie_premiuem/core/utils"
	"net/http"
)

type RegisterUserSerializer struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (s *RegisterUserSerializer) Serialize() (bool, map[string]string) {
	// Validate fields
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		fields := utils.FieldValidator(err)
		if fields != nil {
			return false, fields
		}
		return false, nil
	}

	return true, map[string]string{}
}

// NewRegisterUserSerializer NewRegisterSerializer Method to deserialize the request body into a RegisterSerializer
func NewRegisterUserSerializer(r *http.Request) *RegisterUserSerializer {
	var serializer RegisterUserSerializer
	json.NewDecoder(r.Body).Decode(&serializer)

	return &serializer
}
