package getTask

import (
	"encoding/json"
	"errors"
)

type ErrorResponse struct {
	Error string
}

type Tasks struct {
	Pk string  `json:"pk"`
	Sk  string  `json:"sk"`
	Id string  `json:"id"`
	Status string  `json:"status"`
	Description string  `json:"description"`
	Title string  `json:"title"`
	Date string  `json:"date"`
	Tags []string `json:"tags"`
	Username string  `json:"username"`
	Size float32 `json:"size"`
}

type UserRequest struct {
	Id string `json:"id"`
	Token string `json:"token"`
}

func ParseUnmarshal(obj string) (Tasks, error)  {
	var tasks Tasks
	err := json.Unmarshal([]byte(obj), &tasks)
	if err != nil {
		return Tasks{}, errors.New("cannot unmarshal json")
	}
	return tasks, nil
}