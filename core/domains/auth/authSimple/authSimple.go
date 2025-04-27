package authSimple

import (
	"encoding/json"
	"errors"
)

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserVerify struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Extra_verification string `json:"extra_verification"`
}

type SuccessResponse struct {
	Status int
	Type  string
	Data  string
	
}

type ErrorResponse struct {
	Error string
}

func ParseUnmarshal(obj interface{}) UserVerify {
	var user UserVerify
	json.Unmarshal([]byte(obj.(string)), &user)
	return user
}


func (u *UserAuth) Verify() (bool, error) {
	if len(u.Username) < 3 {
		return false, errors.New("username too short")
	}

	if len(u.Password) < 6 {
		return false, errors.New("password too short")
	}

	return true, nil
}