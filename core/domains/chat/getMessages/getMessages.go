package getMessages

// "github.com/Grs2080w/grp_server/core/domains/chat/getMessage"

import (
	"encoding/json"
	"errors"
)

type Message struct {
	Id string `json:"id"`
	Message string `json:"message"`
	Date string `json:"date"`
	Hour string `json:"hour"`
}

type ErrorResponse struct {
    Error string `json:"error"`
}

func ParseMessages(obj string) ([]Message, error) {
	var messages []Message
	err := json.Unmarshal([]byte(obj), &messages)
	if err != nil {
		return nil, errors.New("error parsing messages")
	}

	return messages, nil
}