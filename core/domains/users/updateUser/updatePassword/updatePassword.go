package updatePassword

import (
	"encoding/json"
	"errors"
)

type UserRequest struct {
	Password string `json:"password"`
	Code     string `json:"code"`
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

	if len(u.Password) < 6 {
		return false, errors.New("password too short")
	}

	if len(u.Code) < 6 {
		return false, errors.New("code too short")
	}

	return true, nil
}