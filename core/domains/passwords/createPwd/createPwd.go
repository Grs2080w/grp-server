package createPwd

// "github.com/Grs2080w/grp_server/core/domains/passwords/createPwd"

import (
	"errors"
)


type Params struct {
	Length int `json:"length"`
	Uppercase_letters bool `json:"uppercase_letters"`
	Lowercase_letters bool `json:"lowercase_letters"`
	Digits bool `json:"digits"`
	Special_characters bool `json:"special_characters"`
}

type Request struct {
	Password string `json:"password"`
	Master string `json:"master"`
	Identifier string `json:"identifier"`
	Mode string `json:"mode"`
	Params Params `json:"params"`
	Tags []string `json:"tags"`
	Size float32 `json:"size"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

/*

{
"Password": "password",
"Master": "master",
"Identifier": "identifier",
"Mode": "auto",
"Params": {
	"Length": 12,
	"Uppercase_letters": true,
	"Lowercase_letters": true,
	"Digits": true,
	"Special_characters": true
},
"Tags": ["tag1", "tag2"],
"Size": 1.5
}

*/

func (r *Request) Validate() error {
	if r.Password == "" && r.Mode == "manual" {
		return errors.New("password is required")
	}

	if r.Master == "" {
		return errors.New("master password is required")
	}

	if len(r.Master) != 8 {
		return errors.New("master password must be 8 characters long")
	}

	if r.Identifier == "" {
		return errors.New("identifier is required")
	}

	if r.Mode == "" {
		return errors.New("mode is required")
	}

	if r.Mode != "manual" && r.Mode != "auto" {
		return errors.New("mode must be either 'manual' or 'auto'")
	}

	if r.Mode == "auto" {
		if r.Params.Length <= 3 {
			return errors.New("length must be greater than 3")
		}
	}

	return nil
}

func ReverseString(s string) string {
	runes := []rune(s)        
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] 
	}
	return string(runes) 
}