package getTags

import (
	"encoding/json"
	"errors"
	"strings"
)

type Tags struct {
	Pk    string                 `dynamodbav:"pk"`
	Sk    string                 `dynamodbav:"sk"`
	Domain string                 `dynamodbav:"domain"`
	Item_id string                 `dynamodbav:"item_id"`
	Tag string                 `dynamodbav:"tag"`
	Date string                 `dynamodbav:"date"`
	Username string  `dynamodbav:"username"`
}

type ErrorResponse struct {
    Error string `json:"error"`
}

func ParseTags(obj string) ([]Tags, error) {
	var tags []Tags
	err := json.Unmarshal([]byte(obj), &tags)
	if err != nil {
		return nil, errors.New("error parsing tags")
	}

	return tags, nil
}

func SplitTagName(tag Tags) string {
	nameTag := strings.Split(tag.Tag, "#")
	return nameTag[1]
}