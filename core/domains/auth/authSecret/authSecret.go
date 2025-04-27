package authSecret

import (
	"encoding/json"
	"errors"
)

type UserAuthSecret struct {
	Token string `json:"token"` 
	Secret_code string `json:"secret_code"`
}

type UserResponse struct {
	Secret_deterministic string `json:"secret_deterministic"`
}

type SuccessResponse struct {
	Status int
	Data  string
}

type ErrorResponse struct {
	Error  string
}

func (u UserAuthSecret) Verify() (bool, error) {
	if u.Token == "" {
		return false, errors.New("token is required")
	}

	if len(u.Secret_code) < 6 {
		return false, errors.New("secret_code is required")
	}

	return true, nil
}

func ParseUnmarshal (obj interface{}) UserResponse {
	var user UserResponse
	json.Unmarshal([]byte(obj.(string)), &user)
	return user
}