package updateEmail

import (
	"encoding/json"
	"errors"
	"net/mail"
)

type UserRequest struct {
	Email string `json:"email"`
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

	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return false, errors.New("invalid email")
	}

	return true, nil
}