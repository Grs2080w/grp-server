package updateTypeVerification

import (
	"encoding/json"
	"errors"
)

type UserRequest struct {
	Type string `json:"type"`
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

	if u.Type != "master_password" && u.Type != "secret_deterministic" {
		return false, errors.New("invalid type verification")
	}

	return true, nil
}