package deleteTask

import (
	"encoding/json"
	"errors"
)

type ErrorResponse struct {
	Error string
}

type SuccessResponse struct {
	Message string
}

type TaskGet struct {
	Date string `json:"date"`
	Tags []string `json:"tags"`
	Size float32 `json:"size"`
}

type UserRequest struct {
	Token string `json:"token"`
	Id string `json:"id"`
}


func ParseUnmarshal(obj string) (TaskGet, error)  {
	var user TaskGet
	err := json.Unmarshal([]byte(obj), &user)
	if err != nil {
		return TaskGet{}, errors.New("cannot unmarshal json")
	}
	return user, nil
}