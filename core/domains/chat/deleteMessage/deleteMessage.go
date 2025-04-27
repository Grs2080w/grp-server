package deleteMessage

// "github.com/Grs2080w/grp_server/core/domains/chat/deleteMessages"

import (
	"encoding/json"
	"errors"
)

type ErrorResponse struct {
    Error string `json:"error"`
}

type SucessResponse struct {
	Message string `json:"message"`
}

type Message struct {
	Pk string  `json:"pk"`
	Sk  string  `json:"sk"`
	Id string `json:"id"`
	Message string `json:"message"`
	Date string `json:"date"`
	Hour string `json:"hour"`
	Username string  `json:"username"`
	Size float32 `json:"size"`
}

func ParseMessages(obj string) (Message, error) {
	var messages Message
	err := json.Unmarshal([]byte(obj), &messages)
	if err != nil {
		return Message{}, errors.New("error parsing messages")
	}

	return messages, nil
}

