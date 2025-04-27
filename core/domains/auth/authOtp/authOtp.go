package authOtp

import (
	"encoding/json"
)

type UserAuthOtp struct {
	Token string `json:"token"` 
}

type UserResponse struct {
	Email string `json:"email"`
}

type SuccessResponse struct {
	Status int
	Data  string
}

type ErrorResponse struct {
	Error  string
}

type Response struct {
	Message string `json:"message"`
}

func ParseUnmarshal (obj interface{}) UserResponse {
	var user UserResponse
	json.Unmarshal([]byte(obj.(string)), &user)
	return user
}