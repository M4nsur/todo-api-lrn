package http

import (
	"encoding/json"
	"fmt"
	"time"
)

type TaskDTO struct {
	Title string
	Description string
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