package addTask

import (
	"errors"
)


type UserRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Tags []string `json:"tags"`
	Size float32 `json:"size"`
}

type ErrorResponse struct {
	Error string
}

type Tasks struct {
	Pk string  `dynamodbav:"pk"`
	Sk  string  `dynamodbav:"sk"`
	Id string  `dynamodbav:"id"`
	Status string  `dynamodbav:"status"`
	Description string  `dynamodbav:"description"`
	Title string  `dynamodbav:"title"`
	Date string  `dynamodbav:"date"`
	Tags []string `dynamodbav:"tags"`
	Username string  `dynamodbav:"username"`
	Size float32 `dynamodbav:"size"`
}


func (u *UserRequest) Verify() (bool, error) {

	if u.Title == ""{
		return false, errors.New("title is required")
	}
	
	if u.Size == 0{
		return false, errors.New("size is required")
	}

	return true, nil
}