package deleteEbook

import (
	"encoding/json"

	model "github.com/Grs2080w/grp_server/core/db/dynamo/ebooks/.model"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

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

func ParseEbook(obj string) (model.Ebook, error) {
	var ebook model.Ebook
	err := json.Unmarshal([]byte(obj), &ebook)
	if err != nil {
		return model.Ebook{}, err
	}
	return ebook, nil
}