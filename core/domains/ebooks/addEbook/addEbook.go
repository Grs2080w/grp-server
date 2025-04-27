package addEbook

import (
	//"github.com/Grs2080w/grp_server/core/db/dynamo/ebooks/.model/model.go"

	"encoding/json"
	"errors"
	str "strconv"
)

type Ebook struct {
	Pk string  `dynamodbav:"pk"`
	Sk  string  `dynamodbav:"sk"`
	Username string  `dynamodbav:"username"`
	Id  string  `dynamodbav:"id"`
    Name string `dynamodbav:"name"` 
	Date string `dynamodbav:"date"`
	Type string `dynamodbav:"type"`
	Tags []string `dynamodbav:"tags"`
	Size float32 `dynamodbav:"size"`
	Extension string `dynamodbav:"extension"`	
}

type Response struct {
	Message string `json:"message"`
	Ebook   Ebook  `json:"ebook"`
}

type Request struct {
	Name    string   `json:"name"`
	Extension string   `json:"extension"`
	Size    float32  `json:"size"`
	Tags    []string `json:"tags"`
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