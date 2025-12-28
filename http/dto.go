package http

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type TaskDTO struct {
	Title       string `json:"title" validate:"required,min=1"`
	Description string `json:"description" validate:"required,min=1"`
}

var validate = validator.New()

func (t TaskDTO) ValidateForCreate() error {
	return validate.Struct(t)
}

type ErrorDTO struct {
	Message string
	Time time.Time
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return fmt.Sprintf(`{"message": "%s"}`, e.Message)
	}
	return string(b)
}