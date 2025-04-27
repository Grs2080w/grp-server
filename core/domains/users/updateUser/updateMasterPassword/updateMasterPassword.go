package updateMasterPassword

import (
	"encoding/json"
	"errors"
)

type UserRequest struct {
	Master_password string `json:"master_password"`
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

	if len(u.Master_password) < 6 {
		return false, errors.New("master password too short")
	}

	return true, nil
}