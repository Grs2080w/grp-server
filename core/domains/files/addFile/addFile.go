package addFile

// "github.com/Grs2080w/grp_server/core/domain/files/addFile"

import (
	"encoding/json"
	"errors"
	str "strconv"

	model "github.com/Grs2080w/grp_server/core/db/dynamo/files/.model"
)


type UserRequest struct {
	Filename string `json:"filename"`
	Type string `json:"type"`
	Size float32 `json:"size"`
	Tags []string `json:"tags"`
}

type FormData struct {}

// type for docs
type Request struct {
	File FormData  `json:"file"`
	Filename string `json:"filename"`
	Type string `json:"type"`
	Size float32 `json:"size"`
	Tags []string `json:"tags"`
}


type Version struct {
	Version string `dynamodbav:"version"`
	Date string `dynamodbav:"date"`
	Id string `dynamodbav:"id"`
	Is_latest bool `dynamodbav:"is_latest"`
	Size float32 `dynamodbav:"size"`
}

type Response struct {
	Message string `json:"message"`
	File   Files `json:"file"`
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

type ErrorResponse struct {
	Error string `json:"error"`
}

func ParseTags(tags string) ([]string, error) {
	var tagList []string
	err := json.Unmarshal([]byte(tags), &tagList)
	if err != nil {
		return nil, errors.New("invalid tags")
	}
	return tagList, nil
}

func ParseSize(size string) (float32, error) {
	size32, err := str.ParseFloat(size, 32)
	if err != nil {
		return 0, errors.New("invalid size")
	}
	return float32(size32), nil
}

func ParseFile(file string) (model.Files, error) {
	var userRequest model.Files
	err := json.Unmarshal([]byte(file), &userRequest)
	if err != nil {
		return model.Files{}, errors.New("invalid file")
	}
	return userRequest, nil
}