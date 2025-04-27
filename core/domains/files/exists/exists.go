package exists

import (
	"encoding/json"
	"errors"
	s "strings"

	model "github.com/Grs2080w/grp_server/core/db/dynamo/files/.model"
)


type UserRequest struct {
	Filename string `json:"filename"`
	Type string `json:"type"`
}

type SuccessResponse struct {
	Exists bool `json:"exists"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func ParseFile(file string) (model.Files, error) {
	var userRequest model.Files
	err := json.Unmarshal([]byte(file), &userRequest)
	if err != nil {
		return model.Files{}, errors.New("invalid file")
	}
	return userRequest, nil
}

func (userRequest *UserRequest) Validate() error {
	if userRequest.Filename == "" || userRequest.Type == "" {
		return errors.New("invalid request")
	}

	if !s.Contains(userRequest.Filename, ".") { return errors.New("invalid filename") }
	return nil
}