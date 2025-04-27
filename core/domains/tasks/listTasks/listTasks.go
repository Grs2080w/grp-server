package listTasks

import (
	"encoding/json"
	"errors"
)

type ErrorResponse struct {
	Error string
}

type UserRequest struct {
	Token string `json:"token"`
}

type Tasks struct {
	Id string  `json:"id"`
	Status string  `json:"status"`
	Description string  `json:"description"`
	Title string  `json:"title"`
}

type TasksList struct {
	Tasks []Tasks `json:"tasks"`
}

func ParseUnmarshal(obj string) ([]Tasks, error)  {
	var tasks []Tasks
	err := json.Unmarshal([]byte(obj), &tasks)
	if err != nil {
		return []Tasks{}, errors.New("cannot unmarshal json")
	}
	return tasks, nil
}