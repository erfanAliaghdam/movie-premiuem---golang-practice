package validator

import (
	"encoding/json"
	"movie_premiuem/core/utils"
	"net/http"
)

type MovieCreateValidator struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	ImageFile   []byte `json:"image_file" validate:"required"`
}

func (v *MovieCreateValidator) Validate() (bool, map[string]string) {
	return utils.ValidateField(v)
}

func NewMovieCreateValidator(r *http.Request) *MovieCreateValidator {
	var validator MovieCreateValidator
	json.NewDecoder(r.Body).Decode(&validator)

	return &validator
}
