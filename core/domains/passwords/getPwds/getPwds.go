package getPwds

import (
	"encoding/json"
)

type ErrorResponse struct {
	Error string `json:"error"`}

type Passwords struct {
	Id string  `dynamodbav:"id"`
	Hash string  `dynamodbav:"hash"`
	Identifier  string  `dynamodbav:"identifier"`
	Tags []string `dynamodbav:"tags"`
	Size float32 `dynamodbav:"size"`
}

func ParsePasswords(obj string) []Passwords {
	var passwords []Passwords
	json.Unmarshal([]byte(obj), &passwords)
	return passwords
}
