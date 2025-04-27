package getlogs

// "github.com/Grs2080w/grp_server/core/domains/logs/getLogs"

import (
	"encoding/json"
)

type Logs struct {
	Pk string  `json:"pk"`
	Sk  string  `json:"sk"`
	Log_id string `json:"log_id"`
	Timestamp int `json:"timestamp"`
	Level string `json:"level"`
	Domain string `json:"domain"`
	Message string `json:"message"`
	Metadata string `json:"metadata"`
	Username string `json:"username"`
	Ip_address string `json:"ip_address"`
	Stack_trace string `json:"stack_trace"`
	Operation string `json:"operation"`
	Status_code string `json:"status_code"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func ParseLogs(data string) ([]Logs, error) {
	var logs []Logs
	err := json.Unmarshal([]byte(data), &logs)
	if err != nil {
		return nil, err
	}
	return logs, nil
}