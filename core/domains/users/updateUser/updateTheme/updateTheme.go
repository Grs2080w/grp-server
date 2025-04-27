package updateTheme

import (
	"encoding/json"
	"errors"
)

type UserRequest struct {
	Theme string `json:"theme"`
}

type UserGet struct {
	Pk string `json:"pk"`
	Sk string `json:"sk"`
}

type SuccessResponse struct {
	Message string
}

type ErrorResponse struct {
	Error string
}

func ParseUnmarshal(data string) (*UserGet, error) {
	var user UserGet
	if err := json.Unmarshal([]byte(data), &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRequest) Verify() (bool, error) {

	if len(u.Theme) < 4 {
		return false, errors.New("theme too short")
	}

	if u.Theme != "dark" && u.Theme != "light" {
		return false, errors.New("invalid theme")
	}

	return true, nil
}