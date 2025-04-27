package getEbooks

// "github.com/Grs2080w/grp_server/core/db/dynamo/ebooks/getEbooks"

import (
	"encoding/json"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type Ebook struct {
	Id  string  `dynamodbav:"id"`
    Name string `dynamodbav:"name"` 
	Date string `dynamodbav:"date"`
	Tags []string `dynamodbav:"tags"`
	Size float32 `dynamodbav:"size"`
	Extension string `dynamodbav:"extension"`	
}


func ParseEbooks(obj string) []Ebook {
	var ebooks []Ebook
	err := json.Unmarshal([]byte(obj), &ebooks)
	if err != nil {
		return []Ebook{}
	}
	return ebooks
}