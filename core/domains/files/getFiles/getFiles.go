package getFiles

import (
	"encoding/json"
	"errors"

	model "github.com/Grs2080w/grp_server/core/db/dynamo/files/.model"
)

// for docs
type UserRequest struct {
	Token string `json:"token"`
}

type SuccessResponse struct {
	Exists bool `json:"exists"`
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

func ParseFiles(file string) ([]model.Files, error) {
	var userRequest []model.Files
	err := json.Unmarshal([]byte(file), &userRequest)
	if err != nil {
		return []model.Files{}, errors.New("invalid file")
	}
	return userRequest, nil
}
