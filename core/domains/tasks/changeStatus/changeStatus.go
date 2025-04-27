package changeStatus

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

type TaskToUpdate struct {
	Pk string  `json:"pk"`
	Sk  string  `json:"sk"`
	Status string  `json:"status"`
}

type UserRequest struct {
	Token string  `json:"token"`
	Id string  `json:"id"` 
}

func ParseUnmarshal(obj string) (TaskToUpdate, error)  {
	var tasks TaskToUpdate
	err := json.Unmarshal([]byte(obj), &tasks)
	if err != nil {
		return TaskToUpdate{}, errors.New("cannot unmarshal json")
	}
	return tasks, nil
}