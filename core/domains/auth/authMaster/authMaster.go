package authMaster

import ("errors"
		"encoding/json"
)

type UserAuthMaster struct {
	Token string `json:"token"`
	Master_password string `json:"master_password"`
}

type UserResponse struct {
	Master_password_hash string `json:"master_password_hash"`
}

type SuccessResponse struct {
	Status int
	Data  string
}

type ErrorResponse struct {
	Error string
}

func (u *UserAuthMaster) Verify() (bool, error) {

	if u.Token == "" {
		return false, errors.New("token is required")
	}
	
	if len(u.Master_password) < 6 {
		return false, errors.New("password too short")
	}

	return true, nil
}

func ParseUnmarshal (obj interface{}) UserResponse {
	var user UserResponse
	json.Unmarshal([]byte(obj.(string)), &user)
	return user
}