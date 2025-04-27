package deletePwd

import (
	"encoding/json"

	models "github.com/Grs2080w/grp_server/core/db/dynamo/passwords/.model"
)

type Passwords struct {
	Id string  `dynamodbav:"id"`
	Hash string  `dynamodbav:"hash"`
	Identifier  string  `dynamodbav:"identifier"`
	Tags []string `dynamodbav:"tags"`
	Size float32 `dynamodbav:"size"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func ParsePassword(obj string) models.Passwords {
	var passwords models.Passwords
	json.Unmarshal([]byte(obj), &passwords)
	return passwords
}