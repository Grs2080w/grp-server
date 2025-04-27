package updateAvatarUrl

import (
	"encoding/json"
	"errors"
)

type UserRequest struct {
	Avatar_url string `json:"avatar_url"`
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
	if len(u.Avatar_url) < 10 {
		return false, errors.New("avatar_url is required")
	}

	return true, nil
}