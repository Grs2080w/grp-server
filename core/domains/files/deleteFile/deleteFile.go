package deleteFile

import (
	"encoding/json"
	"errors"
	s "strings"

	model "github.com/Grs2080w/grp_server/core/db/dynamo/files/.model"
)

type UserRequest struct {
	Filename string `json:"filename"`
	Type string `json:"type"`
	Id string `json:"id"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}


type Version struct {
	Version string `dynamodbav:"version"`
	Date string `dynamodbav:"date"`
	Id string `dynamodbav:"id"`
	Is_latest bool `dynamodbav:"is_latest"`
	Size float32 `dynamodbav:"size"`
}

type Files struct {
	Pk string  `dynamodbav:"pk"`
	Sk  string  `dynamodbav:"sk"`
	Filename string `dynamodbav:"filename"`
    Type string `dynamodbav:"type"`
	Versions []Version `dynamodbav:"versions"`
	Username string  `dynamodbav:"username"`
	Tags []string `dynamodbav:"tags"`
	Extension string `dynamodbav:"extension"`
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
	if userRequest.Filename == "" || userRequest.Type == "" || userRequest.Id == "" {
		return errors.New("invalid request")
	}

	if !s.Contains(userRequest.Filename, ".") { return errors.New("invalid filename") }
	return nil
}