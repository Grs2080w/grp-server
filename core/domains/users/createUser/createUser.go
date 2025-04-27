package createUser

import (
	"errors"
	"net/mail"
)

type Users struct {
	Pk                   string `json:"pk"`
	Sk                   string `json:"sk"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	Email                string `json:"email"`
	Storage_used         string `json:"storage_used"`
	Total_files          string `json:"total_files"`
	Avatar_url           string `json:"avatar_url"`
	Theme_preferences    string `json:"theme_preferences"`
	Language             string `json:"language"`
	Failed_logins        string `json:"failed_logins"`
	Extra_verification   string `json:"extra_verification"`
	Master_password_hash string `json:"master_password_hash"`
	Secret_deterministic string `json:"secret_deterministic"`
}

type UserRequest struct {
	Username          string `json:"username"`
	Password          string `json:"password"`
	Email             string `json:"email"`
	Theme_preferences string `json:"theme_preferences"`
	Language          string `json:"language"`
	Extra_verification string `json:"extra_verification"`
	Code              string `json:"code"`
}

type ErrorResponse struct {
	Error string
}

type SuccessResponse struct {
	Message string `json:"message"`
}

func (u *UserRequest) Verify() (bool, error) {

	_, err := mail.ParseAddress(u.Email) 

	if len(u.Username) < 3 {
		return false, errors.New("username must be at least 3 characters")
	}

	if len(u.Password) < 6 {
		return false, errors.New("password must be at least 8 characters")
	}

	if err != nil {
		return false, errors.New("invalid email")
	}

	if len(u.Theme_preferences) < 4 {
		return false, errors.New("theme_preferences must be at least 4 characters")
	}

	if u.Language != "en" && u.Language != "pt-br" {
		return false, errors.New("language must be en or pt-br")
	}

	if u.Extra_verification != "master_password" && u.Extra_verification != "secret_deterministic" {
		return false, errors.New("invalid type verification")
	}

	if len(u.Code) < 6 && u.Extra_verification == "master_password" {
		return false, errors.New("code must be at least 6 characters")
	}

	if len(u.Code) < 3 && u.Extra_verification == "secret_deterministic" {
		return false, errors.New("code must be at least 3 characters")
	}

	return true, nil

}
