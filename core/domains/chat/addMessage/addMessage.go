package addMessage

// "github.com/Grs2080w/grp_server/core/domains/chat/addMessage"

import (
	"errors"
)

type Request struct {
    Message string `json:"message"`
    Size float32 `json:"size"`
}

type ErrorResponse struct {
    Error string `json:"error"`
}

type SuccessResponse struct {
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

func (r *Request) ParseReq() (error) {
    if r.Message == "" {
        return errors.New("Message do not be empty")
    }

    return nil

}


